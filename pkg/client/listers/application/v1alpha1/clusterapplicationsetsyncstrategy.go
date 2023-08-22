// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterApplicationSetSyncStrategyLister helps list ClusterApplicationSetSyncStrategies.
// All objects returned here must be treated as read-only.
type ClusterApplicationSetSyncStrategyLister interface {
	// List lists all ClusterApplicationSetSyncStrategies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterApplicationSetSyncStrategy, err error)
	// ClusterApplicationSetSyncStrategies returns an object that can list and get ClusterApplicationSetSyncStrategies.
	ClusterApplicationSetSyncStrategies(namespace string) ClusterApplicationSetSyncStrategyNamespaceLister
	ClusterApplicationSetSyncStrategyListerExpansion
}

// clusterApplicationSetSyncStrategyLister implements the ClusterApplicationSetSyncStrategyLister interface.
type clusterApplicationSetSyncStrategyLister struct {
	indexer cache.Indexer
}

// NewClusterApplicationSetSyncStrategyLister returns a new ClusterApplicationSetSyncStrategyLister.
func NewClusterApplicationSetSyncStrategyLister(indexer cache.Indexer) ClusterApplicationSetSyncStrategyLister {
	return &clusterApplicationSetSyncStrategyLister{indexer: indexer}
}

// List lists all ClusterApplicationSetSyncStrategies in the indexer.
func (s *clusterApplicationSetSyncStrategyLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterApplicationSetSyncStrategy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterApplicationSetSyncStrategy))
	})
	return ret, err
}

// ClusterApplicationSetSyncStrategies returns an object that can list and get ClusterApplicationSetSyncStrategies.
func (s *clusterApplicationSetSyncStrategyLister) ClusterApplicationSetSyncStrategies(namespace string) ClusterApplicationSetSyncStrategyNamespaceLister {
	return clusterApplicationSetSyncStrategyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterApplicationSetSyncStrategyNamespaceLister helps list and get ClusterApplicationSetSyncStrategies.
// All objects returned here must be treated as read-only.
type ClusterApplicationSetSyncStrategyNamespaceLister interface {
	// List lists all ClusterApplicationSetSyncStrategies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterApplicationSetSyncStrategy, err error)
	// Get retrieves the ClusterApplicationSetSyncStrategy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterApplicationSetSyncStrategy, error)
	ClusterApplicationSetSyncStrategyNamespaceListerExpansion
}

// clusterApplicationSetSyncStrategyNamespaceLister implements the ClusterApplicationSetSyncStrategyNamespaceLister
// interface.
type clusterApplicationSetSyncStrategyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterApplicationSetSyncStrategies in the indexer for a given namespace.
func (s clusterApplicationSetSyncStrategyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterApplicationSetSyncStrategy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterApplicationSetSyncStrategy))
	})
	return ret, err
}

// Get retrieves the ClusterApplicationSetSyncStrategy from the indexer for a given namespace and name.
func (s clusterApplicationSetSyncStrategyNamespaceLister) Get(name string) (*v1alpha1.ClusterApplicationSetSyncStrategy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterapplicationsetsyncstrategy"), name)
	}
	return obj.(*v1alpha1.ClusterApplicationSetSyncStrategy), nil
}
