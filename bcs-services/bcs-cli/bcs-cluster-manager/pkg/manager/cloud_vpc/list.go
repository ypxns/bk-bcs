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

package cloudvpc

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// List 查询云私有网络列表
func (c *CloudVPCMgr) List(req types.ListCloudVPCReq) (types.ListCloudVPCResp, error) {
	var (
		resp types.ListCloudVPCResp
		err  error
	)

	servResp, err := c.client.ListCloudVPC(c.ctx, &clustermanager.ListCloudVPCRequest{
		NetworkType: req.NetworkType,
	})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.CloudVPC, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, &types.CloudVPC{
			CloudID:     v.CloudID,
			Region:      v.Region,
			RegionName:  v.RegionName,
			NetworkType: v.NetworkType,
			VPCID:       v.VpcID,
			VPCName:     v.VpcName,
			Available:   v.Available,
			Extra:       v.Extra,
			Creator:     v.Creator,
			Updater:     v.Updater,
			CreatTime:   v.CreatTime,
			UpdateTime:  v.UpdateTime,
		})
	}

	return resp, nil
}
