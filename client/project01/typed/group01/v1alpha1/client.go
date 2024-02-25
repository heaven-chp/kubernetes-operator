package v1alpha1

import (
	"kubernetes-operator/api/project01/group01/v1alpha1"
	"kubernetes-operator/client/project01/scheme"
	"net/http"

	"k8s.io/client-go/rest"
)

type Group01V1alpha1Interface interface {
	RESTClient() rest.Interface

	Resource01sGetter
}

type Group01V1alpha1Client struct {
	restClient rest.Interface
}

func (this *Group01V1alpha1Client) RESTClient() rest.Interface {
	if this == nil {
		return nil
	}
	return this.restClient
}

func (this *Group01V1alpha1Client) Resource01s(namespace string) Resource01Interface {
	return newResource01s(this, namespace)
}

func NewForConfig(c *rest.Config) (*Group01V1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

func NewForConfigAndClient(c *rest.Config, h *http.Client) (*Group01V1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &Group01V1alpha1Client{client}, nil
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}
