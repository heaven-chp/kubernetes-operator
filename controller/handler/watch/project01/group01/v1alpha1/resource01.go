package v1alpha1

import (
	"kubernetes-operator/api/project01/group01/v1alpha1"
	"kubernetes-operator/utility"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

type Resource01 struct {
}

func (this *Resource01) OnAdded(object runtime.Object, isInInitialList bool) {
	resource01 := object.(*v1alpha1.Resource01)

	klog.InfoS("OnAdded start", "name", resource01.Name, "namespace", resource01.Namespace, "isInInitialList", isInInitialList)
	defer klog.InfoS("OnAdded end", "name", resource01.Name, "namespace", resource01.Namespace, "isInInitialList", isInInitialList)
}

func (this *Resource01) OnModified(object runtime.Object) {
	resource01 := object.(*v1alpha1.Resource01)

	klog.InfoS("OnModified start", "name", resource01.Name, "namespace", resource01.Namespace)
	defer klog.InfoS("OnModified end", "name", resource01.Name, "namespace", resource01.Namespace)
}

func (this *Resource01) OnDeleted(object runtime.Object) {
	resource01 := object.(*v1alpha1.Resource01)

	klog.InfoS("OnDeleted start", "name", resource01.Name, "namespace", resource01.Namespace)
	defer klog.InfoS("OnDeleted end", "name", resource01.Name, "namespace", resource01.Namespace)
}

func (this *Resource01) OnBookmark(object runtime.Object) {
	resource01 := object.(*v1alpha1.Resource01)

	klog.InfoS("OnBookmark start", "name", resource01.Name, "namespace", resource01.Namespace)
	defer klog.InfoS("OnBookmark end", "name", resource01.Name, "namespace", resource01.Namespace)
}

func (this *Resource01) OnError(object runtime.Object) {
	klog.InfoS("OnError start", "object", object)
	defer klog.InfoS("OnError end", "object", object)
}

func (this *Resource01) GetList(config *rest.Config) (runtime.Object, error) {
	list, err := utility.List[v1alpha1.Resource01List](config, coreV1.NamespaceAll)

	return &list, err
}

func (this *Resource01) GetWatch(config *rest.Config, options metaV1.ListOptions) (watch.Interface, error) {
	return utility.WatchWithOptions[v1alpha1.Resource01](config, coreV1.NamespaceAll, options)
}
