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

package gse

import "testing"

func getNewClient() *Client {
	cli, err := NewGseClient(Options{
		Enable:     true,
		AppCode:    "bcs-xx",
		AppSecret:  "xxx",
		Server:     "xxx",
		BKUserName: "xxx",
		Debug:      true,
	})
	if err != nil {
		return nil
	}

	return cli
}

func TestGetAgentStatus(t *testing.T) {
	cli := getNewClient()

	resp, err := cli.GetAgentStatus(&GetAgentStatusReqV2{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", resp)
}
