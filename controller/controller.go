package controller

import (
	"context"
	informerHandler "kubernetes-operator/controller/handler/informer"
	watchHandler "kubernetes-operator/controller/handler/watch"
	"kubernetes-operator/utility"
	"sync"

	"k8s.io/apimachinery/pkg/api/meta"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	toolsWatch "k8s.io/client-go/tools/watch"
)

type watchHandlerInfo struct {
	handler watchHandler.HandlerInterface
	config  *rest.Config

	list                runtime.Object
	lastResourceVersion string
}

var waitGroup = new(sync.WaitGroup)
var stopper chan struct{}
var retryWatchers []*toolsWatch.RetryWatcher

var informers = map[informerHandler.HandlerInterface]cache.SharedIndexInformer{}
var watchHandlerInfos = map[watchHandler.HandlerInterface]watchHandlerInfo{}

func RegisterHandlerForInformer[Handler informerHandler.HandlerInterface](config *rest.Config) error {
	var handler Handler

	if informer, err := handler.GetInfomer(config); err != nil {
		return err
	} else {
		informer.AddEventHandler(
			cache.ResourceEventHandlerDetailedFuncs{
				AddFunc: func(obj interface{}, isInInitialList bool) {
					handler.OnAdd(obj, isInInitialList)
				},
				UpdateFunc: func(oldObj, newObj interface{}) {
					handler.OnUpdate(oldObj, newObj)
				},
				DeleteFunc: func(obj interface{}) {
					handler.OnDelete(obj)
				},
			})

		informers[handler] = informer

		return nil
	}
}

func UnregisterHandlerForInformer(handler informerHandler.HandlerInterface) {
	delete(informers, handler)
}

func RegisterHandlerForWatch[Handler watchHandler.HandlerInterface](config *rest.Config) error {
	var handler Handler

	untilWithoutRetryFunc := func(event watch.Event) (bool, error) { return true, nil }

	if list, err := handler.GetList(config); err != nil {
		return err
	} else if watcher, err := handler.GetWatch(config, metaV1.ListOptions{}); err != nil {
		return err
	} else if lastEvent, err := toolsWatch.UntilWithoutRetry(context.Background(), watcher, untilWithoutRetryFunc); err != nil {
		return err
	} else if lastResourceVersion, err := utility.ObjectToResourceVersion(lastEvent.Object); err != nil {
		return err
	} else {
		watchHandlerInfos[handler] = watchHandlerInfo{
			handler:             handler,
			config:              config,
			list:                list,
			lastResourceVersion: lastResourceVersion,
		}

		return nil
	}

}

func UnregisterHandlerForWatch(handler watchHandler.HandlerInterface) {
	delete(watchHandlerInfos, handler)
}

func startForInformer(informer cache.SharedIndexInformer) {
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		informer.Run(stopper)
	}()
}

func startForWatch(handlerInfo watchHandlerInfo) error {
	eachListItemFunc := func(object runtime.Object) error {
		handlerInfo.handler.OnAdded(object, true)
		return nil
	}

	if err := meta.EachListItem(handlerInfo.list, eachListItemFunc); err != nil {
		return err
	} else {
		watchFunc := func(options metaV1.ListOptions) (watch.Interface, error) {
			return handlerInfo.handler.GetWatch(handlerInfo.config, options)
		}

		if retryWatcher, err := toolsWatch.NewRetryWatcher(handlerInfo.lastResourceVersion, &cache.ListWatch{WatchFunc: watchFunc}); err != nil {
			return err
		} else {
			retryWatchers = append(retryWatchers, retryWatcher)

			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()

				for event := range retryWatcher.ResultChan() {
					switch event.Type {
					case watch.Added:
						handlerInfo.handler.OnAdded(event.Object, false)
					case watch.Modified:
						handlerInfo.handler.OnModified(event.Object)
					case watch.Deleted:
						handlerInfo.handler.OnDeleted(event.Object)
					case watch.Bookmark:
						handlerInfo.handler.OnBookmark(event.Object)
					case watch.Error:
						handlerInfo.handler.OnError(event.Object)
					}
				}
			}()
		}

		return nil
	}
}

func start() error {
	for _, informer := range informers {
		startForInformer(informer)
	}

	for _, handlerInfo := range watchHandlerInfos {
		if err := startForWatch(handlerInfo); err != nil {
			return err
		}
	}

	return nil
}

func Start() error {
	Stop()

	stopper = make(chan struct{})

	if err := start(); err != nil {
		return err
	}

	return nil
}

func Stop() {
	if stopper != nil {
		close(stopper)
		stopper = nil
	}

	for _, retryWatcher := range retryWatchers {
		retryWatcher.Stop()
	}
	retryWatchers = nil

	waitGroup.Wait()
}
