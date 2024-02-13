# podinfo-custom
A custom Podinfo deployment created with `Kubebuilder` and using the `deploy-image/v1-alpha` plugin. A redis cache is deployed alongside it.

## Todo
- Hook up the controller logic for enabling redis on the custom deployment, currently redis enablement is static
- Expose the Podinfo deployment as a `ClusterIP` service

## Description
To get started:
- Install the CRD on your cluster
- Deploy the controller to the cluster with the specified `IMG` environment variable, if testing locally run:
  ```sh
  make build && make run
  ```
- Create an instance of the `MyAppResource` kind by applying the sample from `config/sample` folder to the cluster

## Getting Started

### Prerequisites
- go version v1.21.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/podinfo-custom:tag
```

**The image has already been published on DockerHub and can be fetched from DockerHub as follows:**
```sh
docker pull docker.io/anandg112/podinfo-custom:latest
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**
```sh
 kustomize build config/crd | kubectl apply -f -
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
export IMG="docker.io/anandg112/podinfo-custom:latest"
cd config/manager && kustomize edit set image controller=${IMG}
kustomize build config/default | kubectl apply -f -
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Sample has been populated with sane default values to test it out

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/podinfo-custom:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/podinfo-custom/<tag or branch>/dist/install.yaml
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024 AnandGautam.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

