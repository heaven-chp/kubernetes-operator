package v1alpha1

import (
	"context"
	"kubernetes-operator/api/project01/group01/v1alpha1"
	"kubernetes-operator/client/project01/scheme"
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type Resource01sGetter interface {
	Resource01s(namespace string) Resource01Interface
}

type Resource01Interface interface {
	Create(ctx context.Context, task *v1alpha1.Resource01, opts metaV1.CreateOptions) (*v1alpha1.Resource01, error)
	Update(ctx context.Context, task *v1alpha1.Resource01, opts metaV1.UpdateOptions) (*v1alpha1.Resource01, error)
	UpdateStatus(ctx context.Context, task *v1alpha1.Resource01, opts metaV1.UpdateOptions) (*v1alpha1.Resource01, error)
	Delete(ctx context.Context, name string, opts metaV1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metaV1.GetOptions) (*v1alpha1.Resource01, error)
	List(ctx context.Context, opts metaV1.ListOptions) (*v1alpha1.Resource01List, error)
	Watch(ctx context.Context, opts metaV1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metaV1.PatchOptions, subresources ...string) (result *v1alpha1.Resource01, err error)
	Resource01Expansion
}

type resource01s struct {
	client    rest.Interface
	namespace string
}

func (this *resource01s) Get(ctx context.Context, name string, options metaV1.GetOptions) (result *v1alpha1.Resource01, err error) {
	result = &v1alpha1.Resource01{}
	err = this.client.Get().
		Namespace(this.namespace).
		Resource("resource01s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (this *resource01s) List(ctx context.Context, opts metaV1.ListOptions) (result *v1alpha1.Resource01List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Resource01List{}
	err = this.client.Get().
		Namespace(this.namespace).
		Resource("resource01s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (this *resource01s) Watch(ctx context.Context, opts metaV1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return this.client.Get().
		Namespace(this.namespace).
		Resource("resource01s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (this *resource01s) Create(ctx context.Context, task *v1alpha1.Resource01, opts metaV1.CreateOptions) (result *v1alpha1.Resource01, err error) {
	result = &v1alpha1.Resource01{}
	err = this.client.Post().
		Namespace(this.namespace).
		Resource("resource01s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(task).
		Do(ctx).
		Into(result)
	return
}

func (this *resource01s) Update(ctx context.Context, task *v1alpha1.Resource01, opts metaV1.UpdateOptions) (result *v1alpha1.Resource01, err error) {
	result = &v1alpha1.Resource01{}
	err = this.client.Put().
		Namespace(this.namespace).
		Resource("resource01s").
		Name(task.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(task).
		Do(ctx).
		Into(result)
	return
}

func (this *resource01s) UpdateStatus(ctx context.Context, task *v1alpha1.Resource01, opts metaV1.UpdateOptions) (result *v1alpha1.Resource01, err error) {
	result = &v1alpha1.Resource01{}
	err = this.client.Put().
		Namespace(this.namespace).
		Resource("resource01s").
		Name(task.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(task).
		Do(ctx).
		Into(result)
	return
}

func (this *resource01s) Delete(ctx context.Context, name string, opts metaV1.DeleteOptions) error {
	return this.client.Delete().
		Namespace(this.namespace).
		Resource("resource01s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (this *resource01s) DeleteCollection(ctx context.Context, opts metaV1.DeleteOptions, listOpts metaV1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return this.client.Delete().
		Namespace(this.namespace).
		Resource("resource01s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (this *resource01s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metaV1.PatchOptions, subresources ...string) (result *v1alpha1.Resource01, err error) {
	result = &v1alpha1.Resource01{}
	err = this.client.Patch(pt).
		Namespace(this.namespace).
		Resource("resource01s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

func newResource01s(client *Group01V1alpha1Client, namespace string) *resource01s {
	return &resource01s{
		client:    client.RESTClient(),
		namespace: namespace,
	}
}
