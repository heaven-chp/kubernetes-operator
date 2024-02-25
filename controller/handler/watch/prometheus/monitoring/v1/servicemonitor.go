package v1

import (
	"kubernetes-operator/utility"

	monitoringV1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

type ServiceMonitor struct {
}

func (this *ServiceMonitor) OnAdded(object runtime.Object, isInInitialList bool) {
	serviceMonitor := object.(*monitoringV1.ServiceMonitor)

	klog.InfoS("OnAdded start", "name", serviceMonitor.Name, "namespace", serviceMonitor.Namespace, "isInInitialList", isInInitialList)
	defer klog.InfoS("OnAdded end", "name", serviceMonitor.Name, "namespace", serviceMonitor.Namespace, "isInInitialList", isInInitialList)
}

func (this *ServiceMonitor) OnModified(object runtime.Object) {
	serviceMonitor := object.(*monitoringV1.ServiceMonitor)

	klog.InfoS("OnModified start", "name", serviceMonitor.Name, "namespace", serviceMonitor.Namespace)
	defer klog.InfoS("OnModified end", "name", serviceMonitor.Name, "namespace", serviceMonitor.Namespace)
}

func (this *ServiceMonitor) OnDeleted(object runtime.Object) {
	serviceMonitor := object.(*monitoringV1.ServiceMonitor)

	klog.InfoS("OnDeleted start", "name", serviceMonitor.Name, "namespace", serviceMonitor.Namespace)
	defer klog.InfoS("OnDeleted end", "name", serviceMonitor.Name, "namespace", serviceMonitor.Namespace)
}

func (this *ServiceMonitor) OnBookmark(object runtime.Object) {
	serviceMonitor := object.(*monitoringV1.ServiceMonitor)

	klog.InfoS("OnBookmark start", "name", serviceMonitor.Name, "namespace", serviceMonitor.Namespace)
	defer klog.InfoS("OnBookmark end", "name", serviceMonitor.Name, "namespace", serviceMonitor.Namespace)
}

func (this *ServiceMonitor) OnError(object runtime.Object) {
	klog.InfoS("OnError start", "object", object)
	defer klog.InfoS("OnError end", "object", object)
}

func (this *ServiceMonitor) GetList(config *rest.Config) (runtime.Object, error) {
	list, err := utility.List[monitoringV1.ServiceMonitorList](config, coreV1.NamespaceAll)

	return &list, err
}

func (this *ServiceMonitor) GetWatch(config *rest.Config, options metaV1.ListOptions) (watch.Interface, error) {
	return utility.WatchWithOptions[monitoringV1.ServiceMonitor](config, coreV1.NamespaceAll, options)
}
