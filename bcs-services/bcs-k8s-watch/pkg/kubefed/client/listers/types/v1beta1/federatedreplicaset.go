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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedReplicaSetLister helps list FederatedReplicaSets.
type FederatedReplicaSetLister interface {
	// List lists all FederatedReplicaSets in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedReplicaSet, err error)
	// FederatedReplicaSets returns an object that can list and get FederatedReplicaSets.
	FederatedReplicaSets(namespace string) FederatedReplicaSetNamespaceLister
	FederatedReplicaSetListerExpansion
}

// federatedReplicaSetLister implements the FederatedReplicaSetLister interface.
type federatedReplicaSetLister struct {
	indexer cache.Indexer
}

// NewFederatedReplicaSetLister returns a new FederatedReplicaSetLister.
func NewFederatedReplicaSetLister(indexer cache.Indexer) FederatedReplicaSetLister {
	return &federatedReplicaSetLister{indexer: indexer}
}

// List lists all FederatedReplicaSets in the indexer.
func (s *federatedReplicaSetLister) List(selector labels.Selector) (ret []*v1beta1.FederatedReplicaSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedReplicaSet))
	})
	return ret, err
}

// FederatedReplicaSets returns an object that can list and get FederatedReplicaSets.
func (s *federatedReplicaSetLister) FederatedReplicaSets(namespace string) FederatedReplicaSetNamespaceLister {
	return federatedReplicaSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedReplicaSetNamespaceLister helps list and get FederatedReplicaSets.
type FederatedReplicaSetNamespaceLister interface {
	// List lists all FederatedReplicaSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedReplicaSet, err error)
	// Get retrieves the FederatedReplicaSet from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedReplicaSet, error)
	FederatedReplicaSetNamespaceListerExpansion
}

// federatedReplicaSetNamespaceLister implements the FederatedReplicaSetNamespaceLister
// interface.
type federatedReplicaSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedReplicaSets in the indexer for a given namespace.
func (s federatedReplicaSetNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedReplicaSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedReplicaSet))
	})
	return ret, err
}

// Get retrieves the FederatedReplicaSet from the indexer for a given namespace and name.
func (s federatedReplicaSetNamespaceLister) Get(name string) (*v1beta1.FederatedReplicaSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedreplicaset"), name)
	}
	return obj.(*v1beta1.FederatedReplicaSet), nil
}
