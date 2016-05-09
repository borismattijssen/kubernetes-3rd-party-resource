package main

import (
	"fmt"
	"net"

	"k8s.io/kubernetes/pkg/api"
	api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/registry/thirdpartyresourcedata"
	"k8s.io/kubernetes/pkg/watch"
)

// host and port for kubie api
const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	thirdPartyClient, err := NewThirdparty(&restclient.Config{
		Host: "http://" + net.JoinHostPort(HOST, PORT),
	})
	if err != nil {
		fmt.Println("Couldn't create 3rd-party client: ", err)
	}
	w, werr := thirdPartyClient.Watch()
	if werr != nil {
		fmt.Println("Couldn't watch workflows: ", werr)
		return
	}
	fmt.Println("Watching..")
	for res := range w.ResultChan() {
		fmt.Println("Something")
		fmt.Println(res)
	}
	return
}

// ThirdPartyClient can be used to access third party resources
type ThirdPartyClient struct {
	*restclient.RESTClient
}

// NewThirdparty creates a new ThirdPartyClient
func NewThirdparty(c *restclient.Config) (*ThirdPartyClient, error) {
	config := *c
	if err := setThirdPartyDefaults(&config); err != nil {
		return nil, err
	}
	client, err := restclient.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ThirdPartyClient{client}, nil
}

// Configuration for RESTClient
func setThirdPartyDefaults(config *restclient.Config) error {
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = restclient.DefaultKubernetesUserAgent()
	}

	// TODO: Group and version shouldn't be hardcoded
	config.GroupVersion = &api_unversioned.GroupVersion{
		Group:   "nerdalize.com",
		Version: "v1alpha1",
	}

	config.Codec = api.Codecs.LegacyCodec(*config.GroupVersion)

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}
	return nil
}

// Watch a workflow using the ThirdPartyClient
func (t *ThirdPartyClient) Watch() (watch.Interface, error) {
	opts := &api.ListOptions{
		TypeMeta: api_unversioned.TypeMeta{
			Kind:       "Workflow",
			APIVersion: "nerdalize.com/v1alpha1",
		},
		Watch: true,
	}

	return t.Get().
		Prefix("watch").
		Namespace("default").
		Resource("workflows").
		VersionedParams(opts, thirdpartyresourcedata.NewThirdPartyParameterCodec(api.ParameterCodec)).
		Watch()
}
