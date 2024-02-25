package informers

import (
	"fmt"
	group01V1alpha1 "kubernetes-operator/api/project01/group01/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
)

type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	case group01V1alpha1.SchemeGroupVersion.WithResource("resource01s"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Group01().V1alpha1().Resource01s().Informer()}, nil
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
