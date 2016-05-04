# Creating a ThirdPartyResource in Kubernetes

Explains how to create a ThirdPartyResource in Kubernetes

## Creating the API endpoint
To create the ThirdPartyResource run the following command:
```
  kubectl create -f 3rd-party-object.yaml
```

In that file `metadata.name` represents `<kind>.<api-group>`. For more information about api groups please see [Kubernetes docs](http://kubernetes.io/docs/api/#api-groups).
Multiple API versions can be specified, <3rd-party-object.yaml> contains only `v1alpha1`.

A new api endpoint is now created, which can be found by navigating to <http://localhost:8080>. The full path for the workflow kind is `/apis/nerdalize.com/v1alpha1/namespaces/<namespace>/workflows`.
Kubernetes now recognizes the `Workflow` kind as specified in <3rd-party-object.yaml>.

## Adding a Workflow
Adding new `Workflow`s can now easily be done using `kubectl`:
```
  kubectl create -f workflow.yaml
```
Note that apiVersion should be specified as `<api-group>/<version>` and kind is `Workflow`.

## More information
Note that at the time of writing most of this functionality is only supported by Kubernetes 1.3.0alpha.

More info can be found at <https://github.com/kubernetes/kubernetes/blob/master/docs/design/extending-api.md>.
