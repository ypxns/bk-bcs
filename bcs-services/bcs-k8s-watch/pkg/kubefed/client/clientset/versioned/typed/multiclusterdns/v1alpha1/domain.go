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

package v1alpha1

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/apis/multiclusterdns/v1alpha1"
	scheme "github.com/Tencent/bk-bcs/bcs-services/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DomainsGetter has a method to return a DomainInterface.
// A group's client should implement this interface.
type DomainsGetter interface {
	Domains(namespace string) DomainInterface
}

// DomainInterface has methods to work with Domain resources.
type DomainInterface interface {
	Create(*v1alpha1.Domain) (*v1alpha1.Domain, error)
	Update(*v1alpha1.Domain) (*v1alpha1.Domain, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Domain, error)
	List(opts v1.ListOptions) (*v1alpha1.DomainList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Domain, err error)
	DomainExpansion
}

// domains implements DomainInterface
type domains struct {
	client rest.Interface
	ns     string
}

// newDomains returns a Domains
func newDomains(c *MulticlusterdnsV1alpha1Client, namespace string) *domains {
	return &domains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the domain, and returns the corresponding domain object, and an error if there is any.
func (c *domains) Get(name string, options v1.GetOptions) (result *v1alpha1.Domain, err error) {
	result = &v1alpha1.Domain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("domains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Domains that match those selectors.
func (c *domains) List(opts v1.ListOptions) (result *v1alpha1.DomainList, err error) {
	result = &v1alpha1.DomainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("domains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested domains.
func (c *domains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("domains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a domain and creates it.  Returns the server's representation of the domain, and an error, if there is any.
func (c *domains) Create(domain *v1alpha1.Domain) (result *v1alpha1.Domain, err error) {
	result = &v1alpha1.Domain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("domains").
		Body(domain).
		Do().
		Into(result)
	return
}

// Update takes the representation of a domain and updates it. Returns the server's representation of the domain, and an error, if there is any.
func (c *domains) Update(domain *v1alpha1.Domain) (result *v1alpha1.Domain, err error) {
	result = &v1alpha1.Domain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("domains").
		Name(domain.Name).
		Body(domain).
		Do().
		Into(result)
	return
}

// Delete takes name of the domain and deletes it. Returns an error if one occurs.
func (c *domains) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("domains").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *domains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("domains").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched domain.
func (c *domains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Domain, err error) {
	result = &v1alpha1.Domain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("domains").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
