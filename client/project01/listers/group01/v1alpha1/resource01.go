package v1alpha1

import (
	"kubernetes-operator/api/project01/group01/v1alpha1"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type Resource01Lister interface {
	List(selector labels.Selector) (ret []*v1alpha1.Resource01, err error)
	Resource01s(namespace string) Resource01NamespaceLister
	Resource01ListerExpansion
}

type resource01Lister struct {
	indexer cache.Indexer
}

func NewResource01Lister(indexer cache.Indexer) Resource01Lister {
	return &resource01Lister{indexer: indexer}
}

func (s *resource01Lister) List(selector labels.Selector) (ret []*v1alpha1.Resource01, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Resource01))
	})
	return ret, err
}

func (s *resource01Lister) Resource01s(namespace string) Resource01NamespaceLister {
	return resource01NamespaceLister{indexer: s.indexer, namespace: namespace}
}

type Resource01NamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.Resource01, err error)
	Get(name string) (*v1alpha1.Resource01, error)
	Resource01NamespaceListerExpansion
}

type resource01NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

func (s resource01NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Resource01, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Resource01))
	})
	return ret, err
}

func (s resource01NamespaceLister) Get(name string) (*v1alpha1.Resource01, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("resource01"), name)
	}
	return obj.(*v1alpha1.Resource01), nil
}
