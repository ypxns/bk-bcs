/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package tasks

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	proto "github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/cloudprovider/qcloud/api"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/common"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/internal/remote/loop"

	"github.com/avast/retry-go"
	as "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/as/v20180419"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
)

// ApplyInstanceMachinesTask update desired nodes task
func ApplyInstanceMachinesTask(taskID string, stepName string) error {
	start := time.Now()

	// get task and task current step
	state, step, err := cloudprovider.GetTaskStateAndCurrentStep(taskID, stepName)
	if err != nil {
		return err
	}
	// previous step successful when retry task
	if step == nil {
		return nil
	}

	// extract parameter && check validate
	clusterID := step.Params[cloudprovider.ClusterIDKey.String()]
	nodeGroupID := step.Params[cloudprovider.NodeGroupIDKey.String()]
	cloudID := step.Params[cloudprovider.CloudIDKey.String()]
	desiredNodes := step.Params[cloudprovider.ScalingNodesNumKey.String()]

	nodeNum, _ := strconv.Atoi(desiredNodes)
	operator := step.Params[cloudprovider.OperatorKey.String()]
	if len(clusterID) == 0 || len(nodeGroupID) == 0 || len(cloudID) == 0 || len(desiredNodes) == 0 || len(operator) == 0 {
		blog.Errorf("ApplyInstanceMachinesTask[%s]: check parameter validate failed", taskID)
		retErr := fmt.Errorf("ApplyInstanceMachinesTask check parameters failed")
		_ = cloudprovider.UpdateNodeGroupDesiredSize(nodeGroupID, nodeNum, true)
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	dependInfo, err := cloudprovider.GetClusterDependBasicInfo(cloudprovider.GetBasicInfoReq{
		ClusterID:   clusterID,
		CloudID:     cloudID,
		NodeGroupID: nodeGroupID,
	})
	if err != nil {
		blog.Errorf("ApplyInstanceMachinesTask[%s]: GetClusterDependBasicInfo failed: %s", taskID, err.Error())
		retErr := fmt.Errorf("ApplyInstanceMachinesTask GetClusterDependBasicInfo failed")
		_ = cloudprovider.UpdateNodeGroupDesiredSize(nodeGroupID, nodeNum, true)
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// inject taskID
	ctx := cloudprovider.WithTaskIDForContext(context.Background(), taskID)
	activity, err := applyInstanceMachines(ctx, dependInfo, uint64(nodeNum))
	if err != nil {
		blog.Errorf("ApplyInstanceMachinesTask[%s]: applyInstanceMachines failed: %s", taskID, err.Error())
		retErr := fmt.Errorf("ApplyInstanceMachinesTask applyInstanceMachines failed %s", err.Error())
		_ = cloudprovider.UpdateNodeGroupDesiredSize(nodeGroupID, nodeNum, true)
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// trans success nodes to cm DB and record common paras, not handle error
	_ = recordClusterInstanceToDB(ctx, activity, state, dependInfo, uint64(nodeNum))

	// update step
	if err := state.UpdateStepSucc(start, stepName); err != nil {
		blog.Errorf("ApplyInstanceMachinesTask[%s] task %s %s update to storage fatal", taskID, taskID, stepName)
		return err
	}
	return nil
}

// applyInstanceMachines apply machines from asg
func applyInstanceMachines(ctx context.Context, info *cloudprovider.CloudDependBasicInfo, nodeNum uint64) (*as.Activity, error) {
	taskID := cloudprovider.GetTaskIDFromContext(ctx)

	var (
		asgID, activityID string
		activity          *as.Activity
		err               error
	)
	asgID, err = getAsgIDByNodePool(ctx, info)
	if err != nil {
		return nil, fmt.Errorf("applyInstanceMachines[%s] getAsgIDByNodePool failed: %v", taskID, err)
	}
	asCli, err := api.NewASClient(info.CmOption)
	if err != nil {
		return nil, err
	}

	// get asgActivityID
	err = loop.LoopDoFunc(context.Background(), func() error {
		activityID, err = asCli.ScaleOutInstances(asgID, nodeNum)
		if err != nil {
			if strings.Contains(err.Error(), as.RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY) {
				blog.Infof("applyInstanceMachines[%s] ScaleOutInstances: %v", taskID, as.RESOURCEUNAVAILABLE_AUTOSCALINGGROUPINACTIVITY)
				return nil
			}
			blog.Errorf("applyInstanceMachines[%s] ScaleOutInstances failed: %v", taskID, err)
			return err
		}
		return loop.EndLoop
	}, loop.LoopInterval(10*time.Second))
	if err != nil || activityID == "" {
		return nil, fmt.Errorf("applyInstanceMachines[%s] ScaleOutInstances failed: %v", taskID, err)
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	// get activityID status
	err = loop.LoopDoFunc(ctx, func() error {
		activity, err = asCli.DescribeAutoScalingActivities(activityID)
		if err != nil {
			blog.Errorf("taskID[%s] DescribeAutoScalingActivities[%s] failed: %v", taskID, activityID, err)
			return nil
		}
		switch *activity.StatusCode {
		case api.SuccessfulActivity.String(), api.SuccessfulPartActivity.String():
			blog.Infof("taskID[%s] DescribeAutoScalingActivities[%s] status[%s]",
				taskID, activityID, *activity.StatusCode)
			return loop.EndLoop
		case api.FailedActivity.String():
			return fmt.Errorf("taskID[%s] DescribeAutoScalingActivities[%s] failed, cause: %v, message: %v",
				taskID, activityID, *activity.Cause, *activity.StatusMessage)
		case api.CancelledActivity.String():
			return fmt.Errorf("taskID[%s] DescribeAutoScalingActivities[%s] failed: %v", taskID, activityID, api.CancelledActivity.String())
		default:
			blog.Infof("taskID[%s] DescribeAutoScalingActivities[%s] still creating, status[%s]",
				taskID, activityID, *activity.StatusCode)
			return nil
		}
	}, loop.LoopInterval(30*time.Second))
	if err != nil {
		blog.Errorf("taskID[%s] applyInstanceMachines[%s] failed: %v", taskID, activityID, err)
		return nil, err
	}

	return activity, nil
}

// recordClusterInstanceToDB already auto build instances to cluster, thus not handle error
func recordClusterInstanceToDB(ctx context.Context, activity *as.Activity, state *cloudprovider.TaskState,
	info *cloudprovider.CloudDependBasicInfo, nodeNum uint64) error {
	// get success instances
	var (
		successInstanceID []string
		failedInstanceID  []string
	)
	taskID := cloudprovider.GetTaskIDFromContext(ctx)

	for _, ins := range activity.ActivityRelatedInstanceSet {
		if *ins.InstanceStatus == api.SuccessfulInstanceAS.String() {
			successInstanceID = append(successInstanceID, *ins.InstanceId)
		} else {
			failedInstanceID = append(failedInstanceID, *ins.InstanceId)
		}
	}
	// rollback desired num
	if len(successInstanceID) != int(nodeNum) {
		_ = cloudprovider.UpdateNodeGroupDesiredSize(info.NodeGroup.NodeGroupID, int(nodeNum)-len(successInstanceID), true)
	}

	// record instanceIDs to task common
	if state.Task.CommonParams == nil {
		state.Task.CommonParams = make(map[string]string)
	}
	if len(successInstanceID) > 0 {
		state.Task.CommonParams[cloudprovider.SuccessNodeIDsKey.String()] = strings.Join(successInstanceID, ",")
		state.Task.CommonParams[cloudprovider.NodeIDsKey.String()] = strings.Join(successInstanceID, ",")
	}
	if len(failedInstanceID) > 0 {
		state.Task.CommonParams[cloudprovider.FailedNodeIDsKey.String()] = strings.Join(failedInstanceID, ",")
	}

	// record successNodes to cluster manager DB
	nodeIPs, err := transInstancesToNode(ctx, successInstanceID, info)
	if err != nil {
		blog.Errorf("recordClusterInstanceToDB[%s] failed: %v", taskID, err)
	}
	if len(nodeIPs) > 0 {
		state.Task.CommonParams[cloudprovider.OriginNodeIPsKey.String()] = strings.Join(nodeIPs, ",")
		state.Task.CommonParams[cloudprovider.NodeIPsKey.String()] = strings.Join(nodeIPs, ",")
	}

	return nil
}

// transInstancesToNode record success nodes to cm DB
func transInstancesToNode(ctx context.Context, successInstanceID []string, info *cloudprovider.CloudDependBasicInfo) ([]string, error) {
	var (
		cvmCli  = api.NodeManager{}
		nodes   = make([]*proto.Node, 0)
		nodeIPs = make([]string, 0)
		err     error
	)

	taskID := cloudprovider.GetTaskIDFromContext(ctx)
	err = retry.Do(func() error {
		nodes, err = cvmCli.ListNodesByInstanceID(successInstanceID, &cloudprovider.ListNodesOption{
			Common:       info.CmOption,
			ClusterVPCID: info.Cluster.VpcID,
		})
		if err != nil {
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		blog.Errorf("transInstancesToNode[%s] failed: %v", taskID, err)
		return nil, err
	}

	for _, n := range nodes {
		nodeIPs = append(nodeIPs, n.InnerIP)
		n.ClusterID = info.NodeGroup.ClusterID
		n.NodeGroupID = info.NodeGroup.NodeGroupID
		n.Passwd = info.NodeGroup.LaunchTemplate.InitLoginPassword
		n.Status = common.StatusInitialization
		err = cloudprovider.SaveNodeInfoToDB(ctx, n, false)
		if err != nil {
			blog.Errorf("transInstancesToNode[%s] SaveNodeInfoToDB[%s] failed: %v", taskID, n.InnerIP, err)
		}
	}

	return nodeIPs, nil
}

func getAsgIDByNodePool(ctx context.Context, info *cloudprovider.CloudDependBasicInfo) (string, error) {
	taskID := cloudprovider.GetTaskIDFromContext(ctx)

	nodePoolID := info.NodeGroup.CloudNodeGroupID
	tkeCli, err := api.NewTkeClient(info.CmOption)
	if err != nil {
		return "", err
	}

	var (
		pool *tke.NodePool
	)
	err = retry.Do(
		func() error {
			pool, err = tkeCli.DescribeClusterNodePoolDetail(info.Cluster.SystemID, nodePoolID)
			if err != nil {
				return err
			}

			return nil
		},
		retry.Context(ctx), retry.Attempts(3),
	)
	if err != nil {
		blog.Errorf("applyInstancesFromNodePool[%s] failed: %v", taskID, err)
		return "", err
	}

	return *pool.AutoscalingGroupId, nil
}

// CheckClusterNodesStatusTask check update desired nodes status task. nodes already add to cluster, thus not rollback desiredNum and only record status
func CheckClusterNodesStatusTask(taskID string, stepName string) error {
	start := time.Now()

	// get task and task current step
	state, step, err := cloudprovider.GetTaskStateAndCurrentStep(taskID, stepName)
	if err != nil {
		return err
	}
	// previous step successful when retry task
	if step == nil {
		return nil
	}

	// step login started here
	// extract parameter && check validate
	clusterID := step.Params[cloudprovider.ClusterIDKey.String()]
	nodeGroupID := step.Params[cloudprovider.NodeGroupIDKey.String()]
	cloudID := step.Params[cloudprovider.CloudIDKey.String()]
	successInstanceID := cloudprovider.ParseNodeIpOrIdFromCommonMap(state.Task.CommonParams,
		cloudprovider.SuccessNodeIDsKey.String(), ",")

	if len(clusterID) == 0 || len(nodeGroupID) == 0 || len(cloudID) == 0 || len(successInstanceID) == 0 {
		blog.Errorf("CheckClusterNodesStatusTask[%s]: check parameter validate failed", taskID)
		retErr := fmt.Errorf("CheckClusterNodesStatusTask check parameters failed")
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}
	dependInfo, err := cloudprovider.GetClusterDependBasicInfo(cloudprovider.GetBasicInfoReq{
		ClusterID:   clusterID,
		CloudID:     cloudID,
		NodeGroupID: nodeGroupID,
	})
	if err != nil {
		blog.Errorf("CheckClusterNodesStatusTask[%s]: GetClusterDependBasicInfo failed: %s", taskID, err.Error())
		retErr := fmt.Errorf("CheckClusterNodesStatusTask GetClusterDependBasicInfo failed")
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// inject taskID
	ctx := cloudprovider.WithTaskIDForContext(context.Background(), taskID)
	successInstances, failureInstances, err := CheckClusterInstanceStatus(ctx, dependInfo, successInstanceID)
	if err != nil || len(successInstances) == 0 {
		// rollback failed nodes

		blog.Errorf("CheckClusterNodesStatusTask[%s]: checkClusterInstanceStatus failed: %s", taskID, err.Error())
		retErr := fmt.Errorf("CheckClusterNodesStatusTask checkClusterInstanceStatus failed")
		_ = state.UpdateStepFailure(start, stepName, retErr)
		return retErr
	}

	// rollback abnormal nodes
	if len(failureInstances) > 0 {
	}

	blog.Infof("CheckClusterNodeStatusTask[%s] delivery succeed[%d] instances[%v] failed[%d] instances[%v]",
		taskID, len(successInstances), successInstances, len(failureInstances), failureInstances)

	// trans instanceIDs to ipList
	ipList := cloudprovider.GetInstanceIPsByID(ctx, successInstances)

	// update response information to task common params
	if state.Task.CommonParams == nil {
		state.Task.CommonParams = make(map[string]string)
	}
	if len(successInstances) > 0 {
		state.Task.CommonParams[cloudprovider.SuccessClusterNodeIDsKey.String()] = strings.Join(successInstances, ",")
	}
	if len(failureInstances) > 0 {
		state.Task.CommonParams[cloudprovider.FailedClusterNodeIDsKey.String()] = strings.Join(failureInstances, ",")
	}

	// successInstance ip list
	if len(ipList) > 0 {
		state.Task.CommonParams[cloudprovider.DynamicNodeIPListKey.String()] = strings.Join(ipList, ",")
		state.Task.CommonParams[cloudprovider.NodeIPsKey.String()] = strings.Join(ipList, ",")
		state.Task.NodeIPList = ipList
	}

	// update step
	if err := state.UpdateStepSucc(start, stepName); err != nil {
		blog.Errorf("CheckClusterNodesStatusTask[%s] task %s %s update to storage fatal", taskID, taskID, stepName)
		return err
	}

	return nil
}
