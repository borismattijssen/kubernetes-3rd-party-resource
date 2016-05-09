package client

import (
	"k8s.io/kubernetes/pkg/api"
	api_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/registry/thirdpartyresourcedata"
	"k8s.io/kubernetes/pkg/watch"
)

// ThirdPartyClient can be used to access third party resources
type ThirdPartyClient struct {
	*restclient.RESTClient
}

type Workflow struct {
	api_unversioned.TypeMeta `json:",inline"`
	api.ObjectMeta           `json:"metadata,omitempty"`

	Spec WorkflowSpec `json:"spec,omitempty"`
}

type WorkflowSpec struct {
	Steps []WorkflowStep `json:"steps"`
}

type WorkflowStep struct {
	Name      string `json:"name"`
	DependsOn string `json:"dependsOn"`
}

type WorkflowList struct {
	api_unversioned.TypeMeta `json:",inline"`
	api_unversioned.ListMeta `json:"metadata,omitempty"`

	Items []Workflow `json:"items"`
}

type WorkflowKind struct{}

func (obj *WorkflowKind) GroupVersionKind() *api_unversioned.GroupVersionKind {
	return api_unversioned.FromAPIVersionAndKind("nerdalize.com/v1alpha1", "WorkflowList")
}

func (obj *WorkflowKind) SetGroupVersionKind(*api_unversioned.GroupVersionKind) {

}

func (*WorkflowList) CodecDecodeSelf() {
	return
}
func (*WorkflowList) CodecEncodeSelf() {
	return
}
func (obj *WorkflowList) GetObjectKind() api_unversioned.ObjectKind {
	return &WorkflowKind{}
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
	config.NegotiatedSerializer = api.Codecs

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

// List all workflows
func (t *ThirdPartyClient) List() (result *WorkflowList, err error) {
	opts := &api.ListOptions{
		TypeMeta: api_unversioned.TypeMeta{
			Kind:       "Workflow",
			APIVersion: "nerdalize.com/v1alpha1",
		},
		Watch: true,
	}
	result = &WorkflowList{}
	err = t.Get().
		Namespace("default").
		Resource("workflows").
		VersionedParams(opts, thirdpartyresourcedata.NewThirdPartyParameterCodec(api.ParameterCodec)).
		Do().
		Into(result)
	return
}
