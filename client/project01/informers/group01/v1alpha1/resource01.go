package v1alpha1

import (
	"context"
	group01V1alpha1 "kubernetes-operator/api/project01/group01/v1alpha1"
	"kubernetes-operator/client/project01/clientset"
	"kubernetes-operator/client/project01/informers/internalinterfaces"
	"kubernetes-operator/client/project01/listers/group01/v1alpha1"
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type Resource01Informer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.Resource01Lister
}

type resource01Informer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func NewFilteredResource01Informer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metaV1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Group01V1alpha1().Resource01s(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metaV1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Group01V1alpha1().Resource01s(namespace).Watch(context.TODO(), options)
			},
		},
		&group01V1alpha1.Resource01{},
		resyncPeriod,
		indexers,
	)
}

func (f *resource01Informer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResource01Informer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resource01Informer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&group01V1alpha1.Resource01{}, f.defaultInformer)
}

func (f *resource01Informer) Lister() v1alpha1.Resource01Lister {
	return v1alpha1.NewResource01Lister(f.Informer().GetIndexer())
}
