package watch

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type HandlerInterface interface {
	OnAdded(object runtime.Object, isInInitialList bool)
	OnModified(object runtime.Object)
	OnDeleted(object runtime.Object)
	OnBookmark(object runtime.Object)
	OnError(object runtime.Object)

	GetList(config *rest.Config) (runtime.Object, error)
	GetWatch(config *rest.Config, options metaV1.ListOptions) (watch.Interface, error)
}
