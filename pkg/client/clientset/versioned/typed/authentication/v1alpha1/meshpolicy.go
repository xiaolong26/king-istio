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

package v1alpha1

import (
	v1alpha1 "kingfisher/king-istio/pkg/apis/authentication/v1alpha1"
	scheme "kingfisher/king-istio/pkg/client/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MeshPoliciesGetter has a method to return a MeshPolicyInterface.
// A group's client should implement this interface.
type MeshPoliciesGetter interface {
	MeshPolicies() MeshPolicyInterface
}

// MeshPolicyInterface has methods to work with MeshPolicy resources.
type MeshPolicyInterface interface {
	Create(*v1alpha1.MeshPolicy) (*v1alpha1.MeshPolicy, error)
	Update(*v1alpha1.MeshPolicy) (*v1alpha1.MeshPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MeshPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.MeshPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MeshPolicy, err error)
	MeshPolicyExpansion
}

// meshPolicies implements MeshPolicyInterface
type meshPolicies struct {
	client rest.Interface
}

// newMeshPolicies returns a MeshPolicies
func newMeshPolicies(c *AuthenticationV1alpha1Client) *meshPolicies {
	return &meshPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the meshPolicy, and returns the corresponding meshPolicy object, and an error if there is any.
func (c *meshPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.MeshPolicy, err error) {
	result = &v1alpha1.MeshPolicy{}
	err = c.client.Get().
		Resource("meshpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MeshPolicies that match those selectors.
func (c *meshPolicies) List(opts v1.ListOptions) (result *v1alpha1.MeshPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MeshPolicyList{}
	err = c.client.Get().
		Resource("meshpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested meshPolicies.
func (c *meshPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("meshpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a meshPolicy and creates it.  Returns the server's representation of the meshPolicy, and an error, if there is any.
func (c *meshPolicies) Create(meshPolicy *v1alpha1.MeshPolicy) (result *v1alpha1.MeshPolicy, err error) {
	result = &v1alpha1.MeshPolicy{}
	err = c.client.Post().
		Resource("meshpolicies").
		Body(meshPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a meshPolicy and updates it. Returns the server's representation of the meshPolicy, and an error, if there is any.
func (c *meshPolicies) Update(meshPolicy *v1alpha1.MeshPolicy) (result *v1alpha1.MeshPolicy, err error) {
	result = &v1alpha1.MeshPolicy{}
	err = c.client.Put().
		Resource("meshpolicies").
		Name(meshPolicy.Name).
		Body(meshPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the meshPolicy and deletes it. Returns an error if one occurs.
func (c *meshPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("meshpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *meshPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("meshpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched meshPolicy.
func (c *meshPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MeshPolicy, err error) {
	result = &v1alpha1.MeshPolicy{}
	err = c.client.Patch(pt).
		Resource("meshpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
