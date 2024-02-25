package v1

import (
	"kubernetes-operator/utility"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

type Pod struct {
}

func (this *Pod) OnAdd(object interface{}, isInInitialList bool) {
	pod := object.(*coreV1.Pod)

	klog.InfoS("OnAdd start", "name", pod.Name, "isInInitialList", isInInitialList)
	defer klog.InfoS("OnAdd end", "name", pod.Name, "isInInitialList", isInInitialList)
}

func (this *Pod) OnUpdate(oldObject, newObject interface{}) {
	_ = oldObject.(*coreV1.Pod)
	newPod := newObject.(*coreV1.Pod)

	klog.InfoS("OnUpdate start", "name", newPod.Name)
	defer klog.InfoS("OnUpdate end", "name", newPod.Name)
}

func (this *Pod) OnDelete(object interface{}) {
	pod := object.(*coreV1.Pod)

	klog.InfoS("OnDelete start", "name", pod.Name)
	defer klog.InfoS("OnDelete end", "name", pod.Name)
}

func (this *Pod) GetInfomer(config *rest.Config) (cache.SharedIndexInformer, error) {
	return utility.Informer[coreV1.Pod](config)
}
