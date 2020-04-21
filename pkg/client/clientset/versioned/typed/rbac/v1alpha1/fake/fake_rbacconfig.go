/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kingfisher/king-istio/pkg/apis/rbac/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRbacConfigs implements RbacConfigInterface
type FakeRbacConfigs struct {
	Fake *FakeRbacV1alpha1
	ns   string
}

var rbacconfigsResource = schema.GroupVersionResource{Group: "rbac.istio.io", Version: "v1alpha1", Resource: "rbacconfigs"}

var rbacconfigsKind = schema.GroupVersionKind{Group: "rbac.istio.io", Version: "v1alpha1", Kind: "RbacConfig"}

// Get takes name of the rbacConfig, and returns the corresponding rbacConfig object, and an error if there is any.
func (c *FakeRbacConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.RbacConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rbacconfigsResource, c.ns, name), &v1alpha1.RbacConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RbacConfig), err
}

// List takes label and field selectors, and returns the list of RbacConfigs that match those selectors.
func (c *FakeRbacConfigs) List(opts v1.ListOptions) (result *v1alpha1.RbacConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rbacconfigsResource, rbacconfigsKind, c.ns, opts), &v1alpha1.RbacConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RbacConfigList{ListMeta: obj.(*v1alpha1.RbacConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.RbacConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rbacConfigs.
func (c *FakeRbacConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rbacconfigsResource, c.ns, opts))

}

// Create takes the representation of a rbacConfig and creates it.  Returns the server's representation of the rbacConfig, and an error, if there is any.
func (c *FakeRbacConfigs) Create(rbacConfig *v1alpha1.RbacConfig) (result *v1alpha1.RbacConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rbacconfigsResource, c.ns, rbacConfig), &v1alpha1.RbacConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RbacConfig), err
}

// Update takes the representation of a rbacConfig and updates it. Returns the server's representation of the rbacConfig, and an error, if there is any.
func (c *FakeRbacConfigs) Update(rbacConfig *v1alpha1.RbacConfig) (result *v1alpha1.RbacConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rbacconfigsResource, c.ns, rbacConfig), &v1alpha1.RbacConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RbacConfig), err
}

// Delete takes name of the rbacConfig and deletes it. Returns an error if one occurs.
func (c *FakeRbacConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rbacconfigsResource, c.ns, name), &v1alpha1.RbacConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRbacConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rbacconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RbacConfigList{})
	return err
}

// Patch applies the patch and returns the patched rbacConfig.
func (c *FakeRbacConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RbacConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rbacconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RbacConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RbacConfig), err
}
