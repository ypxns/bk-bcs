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

package node

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// Drain 节点pod迁移,将节点上的 Pod 驱逐
func (c *NodeMgr) Drain(req types.DrainNodeReq) (types.DrainNodeResp, error) {
	var (
		resp types.DrainNodeResp
		err  error
	)

	servResp, err := c.client.DrainNode(c.ctx, &clustermanager.DrainNodeRequest{
		InnerIPs:  req.InnerIPs,
		ClusterID: req.ClusterID,
		Updater:   "bcs",
	})
	if err != nil {
		return resp, err
	}

	resp.Data = make([]string, 0)

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = servResp.Fail

	return resp, nil
}
