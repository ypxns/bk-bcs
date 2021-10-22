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

package fake

import (
	"context"

	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/apis/tkex/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHookRuns implements HookRunInterface
type FakeHookRuns struct {
	Fake *FakeTkexV1alpha1
	ns   string
}

var hookrunsResource = schema.GroupVersionResource{Group: "tkex", Version: "v1alpha1", Resource: "hookruns"}

var hookrunsKind = schema.GroupVersionKind{Group: "tkex", Version: "v1alpha1", Kind: "HookRun"}

// Get takes name of the hookRun, and returns the corresponding hookRun object, and an error if there is any.
func (c *FakeHookRuns) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HookRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hookrunsResource, c.ns, name), &v1alpha1.HookRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookRun), err
}

// List takes label and field selectors, and returns the list of HookRuns that match those selectors.
func (c *FakeHookRuns) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HookRunList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hookrunsResource, hookrunsKind, c.ns, opts), &v1alpha1.HookRunList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HookRunList{ListMeta: obj.(*v1alpha1.HookRunList).ListMeta}
	for _, item := range obj.(*v1alpha1.HookRunList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hookRuns.
func (c *FakeHookRuns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hookrunsResource, c.ns, opts))

}

// Create takes the representation of a hookRun and creates it.  Returns the server's representation of the hookRun, and an error, if there is any.
func (c *FakeHookRuns) Create(ctx context.Context, hookRun *v1alpha1.HookRun, opts v1.CreateOptions) (result *v1alpha1.HookRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hookrunsResource, c.ns, hookRun), &v1alpha1.HookRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookRun), err
}

// Update takes the representation of a hookRun and updates it. Returns the server's representation of the hookRun, and an error, if there is any.
func (c *FakeHookRuns) Update(ctx context.Context, hookRun *v1alpha1.HookRun, opts v1.UpdateOptions) (result *v1alpha1.HookRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hookrunsResource, c.ns, hookRun), &v1alpha1.HookRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookRun), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHookRuns) UpdateStatus(ctx context.Context, hookRun *v1alpha1.HookRun, opts v1.UpdateOptions) (*v1alpha1.HookRun, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hookrunsResource, "status", c.ns, hookRun), &v1alpha1.HookRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookRun), err
}

// Delete takes name of the hookRun and deletes it. Returns an error if one occurs.
func (c *FakeHookRuns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hookrunsResource, c.ns, name), &v1alpha1.HookRun{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHookRuns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hookrunsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HookRunList{})
	return err
}

// Patch applies the patch and returns the patched hookRun.
func (c *FakeHookRuns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HookRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hookrunsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HookRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookRun), err
}
