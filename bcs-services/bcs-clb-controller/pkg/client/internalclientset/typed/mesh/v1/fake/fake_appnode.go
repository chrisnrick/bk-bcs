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
	mesh_v1 "bk-bcs/bcs-services/bcs-clb-controller/pkg/apis/mesh/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppNodes implements AppNodeInterface
type FakeAppNodes struct {
	Fake *FakeMeshV1
	ns   string
}

var appnodesResource = schema.GroupVersionResource{Group: "mesh.bmsf.tencent.com", Version: "v1", Resource: "appnodes"}

var appnodesKind = schema.GroupVersionKind{Group: "mesh.bmsf.tencent.com", Version: "v1", Kind: "AppNode"}

// Get takes name of the appNode, and returns the corresponding appNode object, and an error if there is any.
func (c *FakeAppNodes) Get(name string, options v1.GetOptions) (result *mesh_v1.AppNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appnodesResource, c.ns, name), &mesh_v1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mesh_v1.AppNode), err
}

// List takes label and field selectors, and returns the list of AppNodes that match those selectors.
func (c *FakeAppNodes) List(opts v1.ListOptions) (result *mesh_v1.AppNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appnodesResource, appnodesKind, c.ns, opts), &mesh_v1.AppNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mesh_v1.AppNodeList{ListMeta: obj.(*mesh_v1.AppNodeList).ListMeta}
	for _, item := range obj.(*mesh_v1.AppNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appNodes.
func (c *FakeAppNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appnodesResource, c.ns, opts))

}

// Create takes the representation of a appNode and creates it.  Returns the server's representation of the appNode, and an error, if there is any.
func (c *FakeAppNodes) Create(appNode *mesh_v1.AppNode) (result *mesh_v1.AppNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appnodesResource, c.ns, appNode), &mesh_v1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mesh_v1.AppNode), err
}

// Update takes the representation of a appNode and updates it. Returns the server's representation of the appNode, and an error, if there is any.
func (c *FakeAppNodes) Update(appNode *mesh_v1.AppNode) (result *mesh_v1.AppNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appnodesResource, c.ns, appNode), &mesh_v1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mesh_v1.AppNode), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppNodes) UpdateStatus(appNode *mesh_v1.AppNode) (*mesh_v1.AppNode, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appnodesResource, "status", c.ns, appNode), &mesh_v1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mesh_v1.AppNode), err
}

// Delete takes name of the appNode and deletes it. Returns an error if one occurs.
func (c *FakeAppNodes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appnodesResource, c.ns, name), &mesh_v1.AppNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appnodesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mesh_v1.AppNodeList{})
	return err
}

// Patch applies the patch and returns the patched appNode.
func (c *FakeAppNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mesh_v1.AppNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appnodesResource, c.ns, name, data, subresources...), &mesh_v1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mesh_v1.AppNode), err
}
