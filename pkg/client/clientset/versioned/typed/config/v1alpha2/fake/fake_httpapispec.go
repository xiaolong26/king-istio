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
	v1alpha2 "kingfisher/king-istio/pkg/apis/config/v1alpha2"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHTTPAPISpecs implements HTTPAPISpecInterface
type FakeHTTPAPISpecs struct {
	Fake *FakeConfigV1alpha2
	ns   string
}

var httpapispecsResource = schema.GroupVersionResource{Group: "config.istio.io", Version: "v1alpha2", Resource: "httpapispecs"}

var httpapispecsKind = schema.GroupVersionKind{Group: "config.istio.io", Version: "v1alpha2", Kind: "HTTPAPISpec"}

// Get takes name of the hTTPAPISpec, and returns the corresponding hTTPAPISpec object, and an error if there is any.
func (c *FakeHTTPAPISpecs) Get(name string, options v1.GetOptions) (result *v1alpha2.HTTPAPISpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httpapispecsResource, c.ns, name), &v1alpha2.HTTPAPISpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HTTPAPISpec), err
}

// List takes label and field selectors, and returns the list of HTTPAPISpecs that match those selectors.
func (c *FakeHTTPAPISpecs) List(opts v1.ListOptions) (result *v1alpha2.HTTPAPISpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httpapispecsResource, httpapispecsKind, c.ns, opts), &v1alpha2.HTTPAPISpecList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.HTTPAPISpecList{ListMeta: obj.(*v1alpha2.HTTPAPISpecList).ListMeta}
	for _, item := range obj.(*v1alpha2.HTTPAPISpecList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hTTPAPISpecs.
func (c *FakeHTTPAPISpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httpapispecsResource, c.ns, opts))

}

// Create takes the representation of a hTTPAPISpec and creates it.  Returns the server's representation of the hTTPAPISpec, and an error, if there is any.
func (c *FakeHTTPAPISpecs) Create(hTTPAPISpec *v1alpha2.HTTPAPISpec) (result *v1alpha2.HTTPAPISpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httpapispecsResource, c.ns, hTTPAPISpec), &v1alpha2.HTTPAPISpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HTTPAPISpec), err
}

// Update takes the representation of a hTTPAPISpec and updates it. Returns the server's representation of the hTTPAPISpec, and an error, if there is any.
func (c *FakeHTTPAPISpecs) Update(hTTPAPISpec *v1alpha2.HTTPAPISpec) (result *v1alpha2.HTTPAPISpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httpapispecsResource, c.ns, hTTPAPISpec), &v1alpha2.HTTPAPISpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HTTPAPISpec), err
}

// Delete takes name of the hTTPAPISpec and deletes it. Returns an error if one occurs.
func (c *FakeHTTPAPISpecs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(httpapispecsResource, c.ns, name), &v1alpha2.HTTPAPISpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHTTPAPISpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httpapispecsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.HTTPAPISpecList{})
	return err
}

// Patch applies the patch and returns the patched hTTPAPISpec.
func (c *FakeHTTPAPISpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.HTTPAPISpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httpapispecsResource, c.ns, name, pt, data, subresources...), &v1alpha2.HTTPAPISpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HTTPAPISpec), err
}
