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

package cloudaccount

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
)

// List 查询云凭证列表
func (c *CloudAccountMgr) List(req types.ListCloudAccountReq) (types.ListCloudAccountResp, error) {
	var (
		resp types.ListCloudAccountResp
		err  error
	)

	servResp, err := c.client.ListCloudAccount(c.ctx, &clustermanager.ListCloudAccountRequest{})
	if err != nil {
		return resp, err
	}

	if servResp != nil && servResp.Code != 0 {
		return resp, errors.New(servResp.Message)
	}

	resp.Data = make([]*types.CloudAccountInfo, 0)

	for _, v := range servResp.Data {
		resp.Data = append(resp.Data, &types.CloudAccountInfo{
			AccountID:   v.Account.AccountID,
			AccountName: v.Account.AccountName,
			ProjectID:   v.Account.ProjectID,
			Desc:        v.Account.Desc,
			Account: types.Account{
				SecretID:          v.Account.Account.SecretID,
				SecretKey:         v.Account.Account.SecretKey,
				SubscriptionID:    v.Account.Account.SubscriptionID,
				TenantID:          v.Account.Account.TenantID,
				ResourceGroupName: v.Account.Account.ResourceGroupName,
				ClientID:          v.Account.Account.ClientID,
				ClientSecret:      v.Account.Account.ClientSecret,
			},
			Clusters: v.Clusters,
		})
	}

	return resp, nil
}
