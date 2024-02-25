package v1alpha1

import (
	"kubernetes-operator/api/project01/group01/v1alpha1"
	"kubernetes-operator/utility"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

type Resource01 struct {
}

func (this *Resource01) OnAdd(object interface{}, isInInitialList bool) {
	resource01 := object.(*v1alpha1.Resource01)

	klog.InfoS("OnAdd start", "name", resource01.Name, "namespace", resource01.Namespace, "isInInitialList", isInInitialList)
	defer klog.InfoS("OnAdd end", "name", resource01.Name, "namespace", resource01.Namespace, "isInInitialList", isInInitialList)
}

func (this *Resource01) OnUpdate(oldObject, newObject interface{}) {
	_ = oldObject.(*v1alpha1.Resource01)
	newResource01 := newObject.(*v1alpha1.Resource01)

	klog.InfoS("OnUpdate start", "name", newResource01.Name, "namespace", newResource01.Namespace)
	defer klog.InfoS("OnUpdate end", "name", newResource01.Name, "namespace", newResource01.Namespace)
}

func (this *Resource01) OnDelete(object interface{}) {
	resource01 := object.(*v1alpha1.Resource01)

	klog.InfoS("OnDelete start", "name", resource01.Name, "namespace", resource01.Namespace)
	defer klog.InfoS("OnDelete end", "name", resource01.Name, "namespace", resource01.Namespace)
}

func (this *Resource01) GetInfomer(config *rest.Config) (cache.SharedIndexInformer, error) {
	return utility.Informer[v1alpha1.Resource01](config)
}
