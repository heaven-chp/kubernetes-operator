package utility

import (
	"context"
	"errors"
	"fmt"
	project01Group01V1alpha1 "kubernetes-operator/api/project01/group01/v1alpha1"
	project01Clientset "kubernetes-operator/client/project01/clientset"
	project01Informers "kubernetes-operator/client/project01/informers"

	monitoringV1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	prometheusInformers "github.com/prometheus-operator/prometheus-operator/pkg/client/informers/externalversions"
	prometheusClient "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	kubernetesInformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type resourceKind interface {
	coreV1.Pod |

		monitoringV1.ServiceMonitor |

		project01Group01V1alpha1.Resource01
}

type resourceListKind interface {
	coreV1.PodList |

		monitoringV1.ServiceMonitorList |

		project01Group01V1alpha1.Resource01List
}

func GetKubernetesClientset(config *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(config)
}

func GetPrometheusClientset(config *rest.Config) (*prometheusClient.Clientset, error) {
	return prometheusClient.NewForConfig(config)
}

func GetProject01Clientset(config *rest.Config) (*project01Clientset.Clientset, error) {
	return project01Clientset.NewForConfig(config)
}

func List[Kind resourceListKind](config *rest.Config, namespace string) (Kind, error) {
	return ListWithOptions[Kind](config, namespace, metaV1.ListOptions{})
}

func ListWithOptions[Kind resourceListKind](config *rest.Config, namespace string, options metaV1.ListOptions) (Kind, error) {
	var resource Kind

	switch p := any(&resource).(type) {
	case *coreV1.PodList:
		if clientset, err := GetKubernetesClientset(config); err != nil {
			return resource, err
		} else if list, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), options); err != nil {
			return resource, err
		} else {
			*p = *list

			return resource, err
		}

	case *monitoringV1.ServiceMonitorList:
		if clientset, err := GetPrometheusClientset(config); err != nil {
			return resource, err
		} else if list, err := clientset.MonitoringV1().ServiceMonitors(namespace).List(context.TODO(), options); err != nil {
			return resource, err
		} else {
			*p = *list

			return resource, err
		}

	case *project01Group01V1alpha1.Resource01List:
		if clientset, err := GetProject01Clientset(config); err != nil {
			return resource, err
		} else if list, err := clientset.Group01V1alpha1().Resource01s(namespace).List(context.TODO(), options); err != nil {
			return resource, err
		} else {
			*p = *list

			return resource, err
		}
	default:
		return resource, errors.New(fmt.Sprintf("not implemented yet - (%T)", resource))
	}
}

func Watch[Kind resourceKind](config *rest.Config, namespace string) (watch.Interface, error) {
	return WatchWithOptions[Kind](config, namespace, metaV1.ListOptions{})
}

func WatchWithOptions[Kind resourceKind](config *rest.Config, namespace string, options metaV1.ListOptions) (watch.Interface, error) {
	var resource Kind

	switch v := any(resource).(type) {
	case coreV1.Pod:
		if clientset, err := GetKubernetesClientset(config); err != nil {
			return nil, err
		} else {
			return clientset.CoreV1().Pods(namespace).Watch(context.TODO(), options)
		}

	case monitoringV1.ServiceMonitor:
		if clientset, err := GetPrometheusClientset(config); err != nil {
			return nil, err
		} else {
			return clientset.MonitoringV1().ServiceMonitors(namespace).Watch(context.TODO(), options)
		}

	case project01Group01V1alpha1.Resource01:
		if clientset, err := GetProject01Clientset(config); err != nil {
			return nil, err
		} else {
			return clientset.Group01V1alpha1().Resource01s(namespace).Watch(context.TODO(), options)
		}
	default:
		return nil, errors.New(fmt.Sprintf("not implemented yet - (%T)", v))
	}
}

func Informer[Kind resourceKind](config *rest.Config) (cache.SharedIndexInformer, error) {
	var resource Kind

	switch v := any(resource).(type) {
	case coreV1.Pod:
		if clientset, err := GetKubernetesClientset(config); err != nil {
			return nil, err
		} else {
			return kubernetesInformers.NewSharedInformerFactory(clientset, 0).Core().V1().Pods().Informer(), nil
		}

	case monitoringV1.ServiceMonitor:
		if clientset, err := GetPrometheusClientset(config); err != nil {
			return nil, err
		} else {
			return prometheusInformers.NewSharedInformerFactory(clientset, 0).Monitoring().V1().ServiceMonitors().Informer(), nil
		}

	case project01Group01V1alpha1.Resource01:
		if clientset, err := GetProject01Clientset(config); err != nil {
			return nil, err
		} else {
			return project01Informers.NewSharedInformerFactory(clientset, 0).Group01().V1alpha1().Resource01s().Informer(), nil
		}
	default:
		return nil, errors.New(fmt.Sprintf("not implemented yet - (%T)", v))
	}
}

func ObjectToMetaObject(object interface{}) (metaV1.Object, error) {
	return meta.Accessor(object)
}

func ObjectToResourceVersion(object interface{}) (string, error) {
	if metaObject, err := ObjectToMetaObject(object); err != nil {
		return "", err
	} else {
		return metaObject.GetResourceVersion(), nil
	}
}
