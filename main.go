package main

import (
	"kubernetes-operator/controller"
	kubernetesCoreV1InformerHandler "kubernetes-operator/controller/handler/informer/kubernetes/core/v1"
	project01Group01V1alpha1InformerHandler "kubernetes-operator/controller/handler/informer/project01/group01/v1alpha1"
	prometheusMonitoringV1InformerHandler "kubernetes-operator/controller/handler/informer/prometheus/monitoring/v1"
	kubernetesCoreV1WatchHandler "kubernetes-operator/controller/handler/watch/kubernetes/core/v1"
	project01Group01V1alpha1WatchHandler "kubernetes-operator/controller/handler/watch/project01/group01/v1alpha1"
	prometheusMonitoringV1WatchHandler "kubernetes-operator/controller/handler/watch/prometheus/monitoring/v1"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

func registerHandler() error {
	if config, err := rest.InClusterConfig(); err != nil {
		return err
	} else if err := controller.RegisterHandlerForInformer[*kubernetesCoreV1InformerHandler.Pod](config); err != nil {
		return err
	} else if err := controller.RegisterHandlerForInformer[*prometheusMonitoringV1InformerHandler.ServiceMonitor](config); err != nil {
		return err
	} else if err := controller.RegisterHandlerForInformer[*project01Group01V1alpha1InformerHandler.Resource01](config); err != nil {
		return err
	} else if err := controller.RegisterHandlerForWatch[*kubernetesCoreV1WatchHandler.Pod](config); err != nil {
		return err
	} else if err := controller.RegisterHandlerForWatch[*prometheusMonitoringV1WatchHandler.ServiceMonitor](config); err != nil {
		return err
	} else if err := controller.RegisterHandlerForWatch[*project01Group01V1alpha1WatchHandler.Resource01](config); err != nil {
		return err
	} else {
		return nil
	}
}

func main() {
	defer klog.Flush()

	klog.InfoS("process start")
	defer klog.InfoS("process end")

	if err := registerHandler(); err != nil {
		klog.ErrorS(err, "")
	} else if err := controller.Start(); err != nil {
		klog.ErrorS(err, "")
	} else {
		defer controller.Stop()
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
