package v1

import (
	"kubernetes-operator/utility"

	monitoringV1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

type ServiceMonitor struct {
}

func (this *ServiceMonitor) OnAdd(object interface{}, isInInitialList bool) {
	serviceMonitor := object.(*monitoringV1.ServiceMonitor)

	klog.InfoS("OnAdd start", "name", serviceMonitor.Name, "isInInitialList", isInInitialList)
	defer klog.InfoS("OnAdd end", "name", serviceMonitor.Name, "isInInitialList", isInInitialList)
}

func (this *ServiceMonitor) OnUpdate(oldObject, newObject interface{}) {
	_ = oldObject.(*monitoringV1.ServiceMonitor)
	newServiceMonitor := newObject.(*monitoringV1.ServiceMonitor)

	klog.InfoS("OnUpdate start", "name", newServiceMonitor.Name)
	defer klog.InfoS("OnUpdate end", "name", newServiceMonitor.Name)
}

func (this *ServiceMonitor) OnDelete(object interface{}) {
	serviceMonitor := object.(*monitoringV1.ServiceMonitor)

	klog.InfoS("OnDelete start", "name", serviceMonitor.Name)
	defer klog.InfoS("OnDelete end", "name", serviceMonitor.Name)
}

func (this *ServiceMonitor) GetInfomer(config *rest.Config) (cache.SharedIndexInformer, error) {
	return utility.Informer[monitoringV1.ServiceMonitor](config)
}
