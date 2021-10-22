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

// FederatedClusterRoleBindingLister helps list FederatedClusterRoleBindings.
type FederatedClusterRoleBindingLister interface {
	// List lists all FederatedClusterRoleBindings in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRoleBinding, err error)
	// FederatedClusterRoleBindings returns an object that can list and get FederatedClusterRoleBindings.
	FederatedClusterRoleBindings(namespace string) FederatedClusterRoleBindingNamespaceLister
	FederatedClusterRoleBindingListerExpansion
}

// federatedClusterRoleBindingLister implements the FederatedClusterRoleBindingLister interface.
type federatedClusterRoleBindingLister struct {
	indexer cache.Indexer
}

// NewFederatedClusterRoleBindingLister returns a new FederatedClusterRoleBindingLister.
func NewFederatedClusterRoleBindingLister(indexer cache.Indexer) FederatedClusterRoleBindingLister {
	return &federatedClusterRoleBindingLister{indexer: indexer}
}

// List lists all FederatedClusterRoleBindings in the indexer.
func (s *federatedClusterRoleBindingLister) List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedClusterRoleBinding))
	})
	return ret, err
}

// FederatedClusterRoleBindings returns an object that can list and get FederatedClusterRoleBindings.
func (s *federatedClusterRoleBindingLister) FederatedClusterRoleBindings(namespace string) FederatedClusterRoleBindingNamespaceLister {
	return federatedClusterRoleBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedClusterRoleBindingNamespaceLister helps list and get FederatedClusterRoleBindings.
type FederatedClusterRoleBindingNamespaceLister interface {
	// List lists all FederatedClusterRoleBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRoleBinding, err error)
	// Get retrieves the FederatedClusterRoleBinding from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedClusterRoleBinding, error)
	FederatedClusterRoleBindingNamespaceListerExpansion
}

// federatedClusterRoleBindingNamespaceLister implements the FederatedClusterRoleBindingNamespaceLister
// interface.
type federatedClusterRoleBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedClusterRoleBindings in the indexer for a given namespace.
func (s federatedClusterRoleBindingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedClusterRoleBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedClusterRoleBinding))
	})
	return ret, err
}

// Get retrieves the FederatedClusterRoleBinding from the indexer for a given namespace and name.
func (s federatedClusterRoleBindingNamespaceLister) Get(name string) (*v1beta1.FederatedClusterRoleBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedclusterrolebinding"), name)
	}
	return obj.(*v1beta1.FederatedClusterRoleBinding), nil
}
