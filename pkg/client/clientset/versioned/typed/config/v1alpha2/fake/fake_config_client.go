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
	v1alpha2 "kingfisher/king-istio/pkg/client/clientset/versioned/typed/config/v1alpha2"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeConfigV1alpha2 struct {
	*testing.Fake
}

func (c *FakeConfigV1alpha2) AttributeManifests(namespace string) v1alpha2.AttributeManifestInterface {
	return &FakeAttributeManifests{c, namespace}
}

func (c *FakeConfigV1alpha2) HTTPAPISpecs(namespace string) v1alpha2.HTTPAPISpecInterface {
	return &FakeHTTPAPISpecs{c, namespace}
}

func (c *FakeConfigV1alpha2) HTTPAPISpecBindings(namespace string) v1alpha2.HTTPAPISpecBindingInterface {
	return &FakeHTTPAPISpecBindings{c, namespace}
}

func (c *FakeConfigV1alpha2) Handlers(namespace string) v1alpha2.HandlerInterface {
	return &FakeHandlers{c, namespace}
}

func (c *FakeConfigV1alpha2) Instances(namespace string) v1alpha2.InstanceInterface {
	return &FakeInstances{c, namespace}
}

func (c *FakeConfigV1alpha2) QuotaSpecs(namespace string) v1alpha2.QuotaSpecInterface {
	return &FakeQuotaSpecs{c, namespace}
}

func (c *FakeConfigV1alpha2) QuotaSpecBindings(namespace string) v1alpha2.QuotaSpecBindingInterface {
	return &FakeQuotaSpecBindings{c, namespace}
}

func (c *FakeConfigV1alpha2) Rules(namespace string) v1alpha2.RuleInterface {
	return &FakeRules{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeConfigV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
