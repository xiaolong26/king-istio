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

package v1alpha2

import (
	v1alpha2 "kingfisher/king-istio/pkg/apis/config/v1alpha2"
	scheme "kingfisher/king-istio/pkg/client/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HTTPAPISpecBindingsGetter has a method to return a HTTPAPISpecBindingInterface.
// A group's client should implement this interface.
type HTTPAPISpecBindingsGetter interface {
	HTTPAPISpecBindings(namespace string) HTTPAPISpecBindingInterface
}

// HTTPAPISpecBindingInterface has methods to work with HTTPAPISpecBinding resources.
type HTTPAPISpecBindingInterface interface {
	Create(*v1alpha2.HTTPAPISpecBinding) (*v1alpha2.HTTPAPISpecBinding, error)
	Update(*v1alpha2.HTTPAPISpecBinding) (*v1alpha2.HTTPAPISpecBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.HTTPAPISpecBinding, error)
	List(opts v1.ListOptions) (*v1alpha2.HTTPAPISpecBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.HTTPAPISpecBinding, err error)
	HTTPAPISpecBindingExpansion
}

// hTTPAPISpecBindings implements HTTPAPISpecBindingInterface
type hTTPAPISpecBindings struct {
	client rest.Interface
	ns     string
}

// newHTTPAPISpecBindings returns a HTTPAPISpecBindings
func newHTTPAPISpecBindings(c *ConfigV1alpha2Client, namespace string) *hTTPAPISpecBindings {
	return &hTTPAPISpecBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hTTPAPISpecBinding, and returns the corresponding hTTPAPISpecBinding object, and an error if there is any.
func (c *hTTPAPISpecBindings) Get(name string, options v1.GetOptions) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	result = &v1alpha2.HTTPAPISpecBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HTTPAPISpecBindings that match those selectors.
func (c *hTTPAPISpecBindings) List(opts v1.ListOptions) (result *v1alpha2.HTTPAPISpecBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.HTTPAPISpecBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hTTPAPISpecBindings.
func (c *hTTPAPISpecBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a hTTPAPISpecBinding and creates it.  Returns the server's representation of the hTTPAPISpecBinding, and an error, if there is any.
func (c *hTTPAPISpecBindings) Create(hTTPAPISpecBinding *v1alpha2.HTTPAPISpecBinding) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	result = &v1alpha2.HTTPAPISpecBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		Body(hTTPAPISpecBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hTTPAPISpecBinding and updates it. Returns the server's representation of the hTTPAPISpecBinding, and an error, if there is any.
func (c *hTTPAPISpecBindings) Update(hTTPAPISpecBinding *v1alpha2.HTTPAPISpecBinding) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	result = &v1alpha2.HTTPAPISpecBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		Name(hTTPAPISpecBinding.Name).
		Body(hTTPAPISpecBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the hTTPAPISpecBinding and deletes it. Returns an error if one occurs.
func (c *hTTPAPISpecBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hTTPAPISpecBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hTTPAPISpecBinding.
func (c *hTTPAPISpecBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	result = &v1alpha2.HTTPAPISpecBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("httpapispecbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
