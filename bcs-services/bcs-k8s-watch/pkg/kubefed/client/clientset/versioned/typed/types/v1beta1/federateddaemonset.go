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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	scheme "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedDaemonSetsGetter has a method to return a FederatedDaemonSetInterface.
// A group's client should implement this interface.
type FederatedDaemonSetsGetter interface {
	FederatedDaemonSets(namespace string) FederatedDaemonSetInterface
}

// FederatedDaemonSetInterface has methods to work with FederatedDaemonSet resources.
type FederatedDaemonSetInterface interface {
	Create(*v1beta1.FederatedDaemonSet) (*v1beta1.FederatedDaemonSet, error)
	Update(*v1beta1.FederatedDaemonSet) (*v1beta1.FederatedDaemonSet, error)
	UpdateStatus(*v1beta1.FederatedDaemonSet) (*v1beta1.FederatedDaemonSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedDaemonSet, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedDaemonSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedDaemonSet, err error)
	FederatedDaemonSetExpansion
}

// federatedDaemonSets implements FederatedDaemonSetInterface
type federatedDaemonSets struct {
	client rest.Interface
	ns     string
}

// newFederatedDaemonSets returns a FederatedDaemonSets
func newFederatedDaemonSets(c *TypesV1beta1Client, namespace string) *federatedDaemonSets {
	return &federatedDaemonSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedDaemonSet, and returns the corresponding federatedDaemonSet object, and an error if there is any.
func (c *federatedDaemonSets) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedDaemonSet, err error) {
	result = &v1beta1.FederatedDaemonSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federateddaemonsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedDaemonSets that match those selectors.
func (c *federatedDaemonSets) List(opts v1.ListOptions) (result *v1beta1.FederatedDaemonSetList, err error) {
	result = &v1beta1.FederatedDaemonSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federateddaemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedDaemonSets.
func (c *federatedDaemonSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federateddaemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedDaemonSet and creates it.  Returns the server's representation of the federatedDaemonSet, and an error, if there is any.
func (c *federatedDaemonSets) Create(federatedDaemonSet *v1beta1.FederatedDaemonSet) (result *v1beta1.FederatedDaemonSet, err error) {
	result = &v1beta1.FederatedDaemonSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federateddaemonsets").
		Body(federatedDaemonSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedDaemonSet and updates it. Returns the server's representation of the federatedDaemonSet, and an error, if there is any.
func (c *federatedDaemonSets) Update(federatedDaemonSet *v1beta1.FederatedDaemonSet) (result *v1beta1.FederatedDaemonSet, err error) {
	result = &v1beta1.FederatedDaemonSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federateddaemonsets").
		Name(federatedDaemonSet.Name).
		Body(federatedDaemonSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedDaemonSets) UpdateStatus(federatedDaemonSet *v1beta1.FederatedDaemonSet) (result *v1beta1.FederatedDaemonSet, err error) {
	result = &v1beta1.FederatedDaemonSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federateddaemonsets").
		Name(federatedDaemonSet.Name).
		SubResource("status").
		Body(federatedDaemonSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedDaemonSet and deletes it. Returns an error if one occurs.
func (c *federatedDaemonSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federateddaemonsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedDaemonSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federateddaemonsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedDaemonSet.
func (c *federatedDaemonSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedDaemonSet, err error) {
	result = &v1beta1.FederatedDaemonSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federateddaemonsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
