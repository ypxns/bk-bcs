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

package v1

import (
	"context"
	"time"

	v1 "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/apis/cloud/v1"
	scheme "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudIPQuotasGetter has a method to return a CloudIPQuotaInterface.
// A group's client should implement this interface.
type CloudIPQuotasGetter interface {
	CloudIPQuotas(namespace string) CloudIPQuotaInterface
}

// CloudIPQuotaInterface has methods to work with CloudIPQuota resources.
type CloudIPQuotaInterface interface {
	Create(ctx context.Context, cloudIPQuota *v1.CloudIPQuota, opts metav1.CreateOptions) (*v1.CloudIPQuota, error)
	Update(ctx context.Context, cloudIPQuota *v1.CloudIPQuota, opts metav1.UpdateOptions) (*v1.CloudIPQuota, error)
	UpdateStatus(ctx context.Context, cloudIPQuota *v1.CloudIPQuota, opts metav1.UpdateOptions) (*v1.CloudIPQuota, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CloudIPQuota, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CloudIPQuotaList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CloudIPQuota, err error)
	CloudIPQuotaExpansion
}

// cloudIPQuotas implements CloudIPQuotaInterface
type cloudIPQuotas struct {
	client rest.Interface
	ns     string
}

// newCloudIPQuotas returns a CloudIPQuotas
func newCloudIPQuotas(c *CloudV1Client, namespace string) *cloudIPQuotas {
	return &cloudIPQuotas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudIPQuota, and returns the corresponding cloudIPQuota object, and an error if there is any.
func (c *cloudIPQuotas) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CloudIPQuota, err error) {
	result = &v1.CloudIPQuota{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudipquotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudIPQuotas that match those selectors.
func (c *cloudIPQuotas) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CloudIPQuotaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CloudIPQuotaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudipquotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudIPQuotas.
func (c *cloudIPQuotas) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudipquotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cloudIPQuota and creates it.  Returns the server's representation of the cloudIPQuota, and an error, if there is any.
func (c *cloudIPQuotas) Create(ctx context.Context, cloudIPQuota *v1.CloudIPQuota, opts metav1.CreateOptions) (result *v1.CloudIPQuota, err error) {
	result = &v1.CloudIPQuota{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudipquotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudIPQuota).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cloudIPQuota and updates it. Returns the server's representation of the cloudIPQuota, and an error, if there is any.
func (c *cloudIPQuotas) Update(ctx context.Context, cloudIPQuota *v1.CloudIPQuota, opts metav1.UpdateOptions) (result *v1.CloudIPQuota, err error) {
	result = &v1.CloudIPQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudipquotas").
		Name(cloudIPQuota.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudIPQuota).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cloudIPQuotas) UpdateStatus(ctx context.Context, cloudIPQuota *v1.CloudIPQuota, opts metav1.UpdateOptions) (result *v1.CloudIPQuota, err error) {
	result = &v1.CloudIPQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudipquotas").
		Name(cloudIPQuota.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudIPQuota).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cloudIPQuota and deletes it. Returns an error if one occurs.
func (c *cloudIPQuotas) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudipquotas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudIPQuotas) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudipquotas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cloudIPQuota.
func (c *cloudIPQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CloudIPQuota, err error) {
	result = &v1.CloudIPQuota{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudipquotas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
