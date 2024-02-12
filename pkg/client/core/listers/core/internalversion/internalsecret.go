// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	core "github.com/gardener/gardener/pkg/apis/core"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InternalSecretLister helps list InternalSecrets.
// All objects returned here must be treated as read-only.
type InternalSecretLister interface {
	// List lists all InternalSecrets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*core.InternalSecret, err error)
	// InternalSecrets returns an object that can list and get InternalSecrets.
	InternalSecrets(namespace string) InternalSecretNamespaceLister
	InternalSecretListerExpansion
}

// internalSecretLister implements the InternalSecretLister interface.
type internalSecretLister struct {
	indexer cache.Indexer
}

// NewInternalSecretLister returns a new InternalSecretLister.
func NewInternalSecretLister(indexer cache.Indexer) InternalSecretLister {
	return &internalSecretLister{indexer: indexer}
}

// List lists all InternalSecrets in the indexer.
func (s *internalSecretLister) List(selector labels.Selector) (ret []*core.InternalSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*core.InternalSecret))
	})
	return ret, err
}

// InternalSecrets returns an object that can list and get InternalSecrets.
func (s *internalSecretLister) InternalSecrets(namespace string) InternalSecretNamespaceLister {
	return internalSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InternalSecretNamespaceLister helps list and get InternalSecrets.
// All objects returned here must be treated as read-only.
type InternalSecretNamespaceLister interface {
	// List lists all InternalSecrets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*core.InternalSecret, err error)
	// Get retrieves the InternalSecret from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*core.InternalSecret, error)
	InternalSecretNamespaceListerExpansion
}

// internalSecretNamespaceLister implements the InternalSecretNamespaceLister
// interface.
type internalSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InternalSecrets in the indexer for a given namespace.
func (s internalSecretNamespaceLister) List(selector labels.Selector) (ret []*core.InternalSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*core.InternalSecret))
	})
	return ret, err
}

// Get retrieves the InternalSecret from the indexer for a given namespace and name.
func (s internalSecretNamespaceLister) Get(name string) (*core.InternalSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(core.Resource("internalsecret"), name)
	}
	return obj.(*core.InternalSecret), nil
}
