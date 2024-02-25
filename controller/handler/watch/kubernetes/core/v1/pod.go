package v1

import (
	"kubernetes-operator/utility"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

type Pod struct {
}

func (this *Pod) OnAdded(object runtime.Object, isInInitialList bool) {
	pod := object.(*coreV1.Pod)

	klog.InfoS("OnAdded start", "name", pod.Name, "namespace", pod.Namespace, "isInInitialList", isInInitialList)
	defer klog.InfoS("OnAdded end", "name", pod.Name, "namespace", pod.Namespace, "isInInitialList", isInInitialList)
}

func (this *Pod) OnModified(object runtime.Object) {
	pod := object.(*coreV1.Pod)

	klog.InfoS("OnModified start", "name", pod.Name, "namespace", pod.Namespace)
	defer klog.InfoS("OnModified end", "name", pod.Name, "namespace", pod.Namespace)
}

func (this *Pod) OnDeleted(object runtime.Object) {
	pod := object.(*coreV1.Pod)

	klog.InfoS("OnDeleted start", "name", pod.Name, "namespace", pod.Namespace)
	defer klog.InfoS("OnDeleted end", "name", pod.Name, "namespace", pod.Namespace)
}

func (this *Pod) OnBookmark(object runtime.Object) {
	pod := object.(*coreV1.Pod)

	klog.InfoS("OnBookmark start", "name", pod.Name, "namespace", pod.Namespace)
	defer klog.InfoS("OnBookmark end", "name", pod.Name, "namespace", pod.Namespace)
}

func (this *Pod) OnError(object runtime.Object) {
	klog.InfoS("OnError start", "object", object)
	defer klog.InfoS("OnError end", "object", object)
}

func (this *Pod) GetList(config *rest.Config) (runtime.Object, error) {
	list, err := utility.List[coreV1.PodList](config, coreV1.NamespaceAll)

	return &list, err
}

func (this *Pod) GetWatch(config *rest.Config, options metaV1.ListOptions) (watch.Interface, error) {
	return utility.WatchWithOptions[coreV1.Pod](config, coreV1.NamespaceAll, options)
}
