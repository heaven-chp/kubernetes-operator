package v1alpha1

import "kubernetes-operator/client/project01/informers/internalinterfaces"

type Interface interface {
	Resource01s() Resource01Informer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

func (v *version) Resource01s() Resource01Informer {
	return &resource01Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
