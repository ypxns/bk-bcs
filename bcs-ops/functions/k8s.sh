#!/bin/bash

#######################################
# Tencent is pleased to support the open source community by making Blueking Container Service available.
# Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
# Licensed under the MIT License (the "License"); you may not use this file except
# in compliance with the License. You may obtain a copy of the License at
# http://opensource.org/licenses/MIT
# Unless required by applicable law or agreed to in writing, software distributed under
# the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
# either express or implied. See the License for the specific language governing permissions and
# limitations under the License.
#######################################

# generic k8s function
# depend on utils.sh
# independent of business

#######################################
# add helmrepo safely
# Arguments:
# $1: repo_name
# $2: repo_url
# Return:
# can't find helm - return 0
# helm update success - return 0
# helm update fail - return 1
#######################################
k8s::safe_add_helmrepo() {
  if ! command -v helm &>/dev/null; then
    utils::log "WARN" "Did helm installed?"
    return 1
  fi

  local repo_name repo_url
  repo_name=$1
  repo_url=$2
  if helm repo list | grep -q "$repo_name"; then
    echo "remove old helm repo: $repo_name"
    helm repo remove "$repo_name"
  fi
  helm repo add "$repo_name" "$repo_url"
  helm repo list
  if ! helm repo update; then
    utils::log "ERROR" "can't update helm repo"
    return 1
  fi
  return 0
}
