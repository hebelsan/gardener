// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SecretBindingLister helps list SecretBindings.
// All objects returned here must be treated as read-only.
type SecretBindingLister interface {
	// List lists all SecretBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.SecretBinding, err error)
	// SecretBindings returns an object that can list and get SecretBindings.
	SecretBindings(namespace string) SecretBindingNamespaceLister
	SecretBindingListerExpansion
}

// secretBindingLister implements the SecretBindingLister interface.
type secretBindingLister struct {
	indexer cache.Indexer
}

// NewSecretBindingLister returns a new SecretBindingLister.
func NewSecretBindingLister(indexer cache.Indexer) SecretBindingLister {
	return &secretBindingLister{indexer: indexer}
}

// List lists all SecretBindings in the indexer.
func (s *secretBindingLister) List(selector labels.Selector) (ret []*v1beta1.SecretBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SecretBinding))
	})
	return ret, err
}

// SecretBindings returns an object that can list and get SecretBindings.
func (s *secretBindingLister) SecretBindings(namespace string) SecretBindingNamespaceLister {
	return secretBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretBindingNamespaceLister helps list and get SecretBindings.
// All objects returned here must be treated as read-only.
type SecretBindingNamespaceLister interface {
	// List lists all SecretBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.SecretBinding, err error)
	// Get retrieves the SecretBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.SecretBinding, error)
	SecretBindingNamespaceListerExpansion
}

// secretBindingNamespaceLister implements the SecretBindingNamespaceLister
// interface.
type secretBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecretBindings in the indexer for a given namespace.
func (s secretBindingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.SecretBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SecretBinding))
	})
	return ret, err
}

// Get retrieves the SecretBinding from the indexer for a given namespace and name.
func (s secretBindingNamespaceLister) Get(name string) (*v1beta1.SecretBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("secretbinding"), name)
	}
	return obj.(*v1beta1.SecretBinding), nil
}
