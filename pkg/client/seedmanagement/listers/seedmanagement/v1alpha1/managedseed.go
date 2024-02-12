// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/gardener/pkg/apis/seedmanagement/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagedSeedLister helps list ManagedSeeds.
// All objects returned here must be treated as read-only.
type ManagedSeedLister interface {
	// List lists all ManagedSeeds in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedSeed, err error)
	// ManagedSeeds returns an object that can list and get ManagedSeeds.
	ManagedSeeds(namespace string) ManagedSeedNamespaceLister
	ManagedSeedListerExpansion
}

// managedSeedLister implements the ManagedSeedLister interface.
type managedSeedLister struct {
	indexer cache.Indexer
}

// NewManagedSeedLister returns a new ManagedSeedLister.
func NewManagedSeedLister(indexer cache.Indexer) ManagedSeedLister {
	return &managedSeedLister{indexer: indexer}
}

// List lists all ManagedSeeds in the indexer.
func (s *managedSeedLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedSeed, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedSeed))
	})
	return ret, err
}

// ManagedSeeds returns an object that can list and get ManagedSeeds.
func (s *managedSeedLister) ManagedSeeds(namespace string) ManagedSeedNamespaceLister {
	return managedSeedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagedSeedNamespaceLister helps list and get ManagedSeeds.
// All objects returned here must be treated as read-only.
type ManagedSeedNamespaceLister interface {
	// List lists all ManagedSeeds in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedSeed, err error)
	// Get retrieves the ManagedSeed from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagedSeed, error)
	ManagedSeedNamespaceListerExpansion
}

// managedSeedNamespaceLister implements the ManagedSeedNamespaceLister
// interface.
type managedSeedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagedSeeds in the indexer for a given namespace.
func (s managedSeedNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedSeed, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedSeed))
	})
	return ret, err
}

// Get retrieves the ManagedSeed from the indexer for a given namespace and name.
func (s managedSeedNamespaceLister) Get(name string) (*v1alpha1.ManagedSeed, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managedseed"), name)
	}
	return obj.(*v1alpha1.ManagedSeed), nil
}
