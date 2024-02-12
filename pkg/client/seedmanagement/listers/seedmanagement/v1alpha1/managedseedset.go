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

// ManagedSeedSetLister helps list ManagedSeedSets.
// All objects returned here must be treated as read-only.
type ManagedSeedSetLister interface {
	// List lists all ManagedSeedSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedSeedSet, err error)
	// ManagedSeedSets returns an object that can list and get ManagedSeedSets.
	ManagedSeedSets(namespace string) ManagedSeedSetNamespaceLister
	ManagedSeedSetListerExpansion
}

// managedSeedSetLister implements the ManagedSeedSetLister interface.
type managedSeedSetLister struct {
	indexer cache.Indexer
}

// NewManagedSeedSetLister returns a new ManagedSeedSetLister.
func NewManagedSeedSetLister(indexer cache.Indexer) ManagedSeedSetLister {
	return &managedSeedSetLister{indexer: indexer}
}

// List lists all ManagedSeedSets in the indexer.
func (s *managedSeedSetLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedSeedSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedSeedSet))
	})
	return ret, err
}

// ManagedSeedSets returns an object that can list and get ManagedSeedSets.
func (s *managedSeedSetLister) ManagedSeedSets(namespace string) ManagedSeedSetNamespaceLister {
	return managedSeedSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagedSeedSetNamespaceLister helps list and get ManagedSeedSets.
// All objects returned here must be treated as read-only.
type ManagedSeedSetNamespaceLister interface {
	// List lists all ManagedSeedSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedSeedSet, err error)
	// Get retrieves the ManagedSeedSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagedSeedSet, error)
	ManagedSeedSetNamespaceListerExpansion
}

// managedSeedSetNamespaceLister implements the ManagedSeedSetNamespaceLister
// interface.
type managedSeedSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagedSeedSets in the indexer for a given namespace.
func (s managedSeedSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedSeedSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedSeedSet))
	})
	return ret, err
}

// Get retrieves the ManagedSeedSet from the indexer for a given namespace and name.
func (s managedSeedSetNamespaceLister) Get(name string) (*v1alpha1.ManagedSeedSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managedseedset"), name)
	}
	return obj.(*v1alpha1.ManagedSeedSet), nil
}
