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

package clean

import (
	"context"
	"fmt"

	nodegroup "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node_group"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	cleanNodesInGroupExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager clean nodesInGroup --clusterID xxx --nodeGroupID xxx --nodes xxx`))
)

func newCleanNodesInGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodesInGroup",
		Short:   "clean nodes to group from bcs-cluster-manager",
		Example: cleanNodesInGroupExample,
		Run:     cleanNodesInGroup,
	}

	cmd.Flags().StringVarP(&clusterID, "clusterID", "c", "", "cluster ID")
	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	cmd.Flags().StringSliceVarP(&nodes, "nodes", "i", []string{}, "node inner ip, for example: -i xxx.xxx.xxx.xxx -i xxx.xxx.xxx.xxx")

	return cmd
}

func cleanNodesInGroup(cmd *cobra.Command, args []string) {
	resp, err := nodegroup.New(context.Background()).CleanNodes(types.CleanNodeGroupReq{
		ClusterID:   clusterID,
		NodeGroupID: nodeGroupID,
		Nodes:       nodes,
	})
	if err != nil {
		klog.Fatalf("clean nodes to group failed: %v", err)
	}

	fmt.Printf("clean nodes to group succeed: taskID: %v\n", resp.TaskID)
}
