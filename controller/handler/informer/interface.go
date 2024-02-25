package informer

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type HandlerInterface interface {
	cache.ResourceEventHandler

	GetInfomer(config *rest.Config) (cache.SharedIndexInformer, error)
}
