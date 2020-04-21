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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "kingfisher/king-istio/pkg/apis/config/v1alpha2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// QuotaSpecLister helps list QuotaSpecs.
type QuotaSpecLister interface {
	// List lists all QuotaSpecs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.QuotaSpec, err error)
	// QuotaSpecs returns an object that can list and get QuotaSpecs.
	QuotaSpecs(namespace string) QuotaSpecNamespaceLister
	QuotaSpecListerExpansion
}

// quotaSpecLister implements the QuotaSpecLister interface.
type quotaSpecLister struct {
	indexer cache.Indexer
}

// NewQuotaSpecLister returns a new QuotaSpecLister.
func NewQuotaSpecLister(indexer cache.Indexer) QuotaSpecLister {
	return &quotaSpecLister{indexer: indexer}
}

// List lists all QuotaSpecs in the indexer.
func (s *quotaSpecLister) List(selector labels.Selector) (ret []*v1alpha2.QuotaSpec, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.QuotaSpec))
	})
	return ret, err
}

// QuotaSpecs returns an object that can list and get QuotaSpecs.
func (s *quotaSpecLister) QuotaSpecs(namespace string) QuotaSpecNamespaceLister {
	return quotaSpecNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// QuotaSpecNamespaceLister helps list and get QuotaSpecs.
type QuotaSpecNamespaceLister interface {
	// List lists all QuotaSpecs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.QuotaSpec, err error)
	// Get retrieves the QuotaSpec from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.QuotaSpec, error)
	QuotaSpecNamespaceListerExpansion
}

// quotaSpecNamespaceLister implements the QuotaSpecNamespaceLister
// interface.
type quotaSpecNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all QuotaSpecs in the indexer for a given namespace.
func (s quotaSpecNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.QuotaSpec, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.QuotaSpec))
	})
	return ret, err
}

// Get retrieves the QuotaSpec from the indexer for a given namespace and name.
func (s quotaSpecNamespaceLister) Get(name string) (*v1alpha2.QuotaSpec, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("quotaspec"), name)
	}
	return obj.(*v1alpha2.QuotaSpec), nil
}
