# Getting Started with Dracon 

This guide will help to quickly setup Dracon on a Kubernetes cluster and get a pipeline running.
The first step is to create a dev Kubernetes cluster in order to deploy Tekton. We suggest you use
KiND to provision a local test cluster quickly. If you already have a K8s cluster then you can skip
directly to the 

## Setting up a [KinD](https://kind.sigs.k8s.io/) cluster

KinD is is a tool for running local Kubernetes clusters using Docker container “nodes”.

1. Create a KinD cluster named `dracon-demo` with its own Docker registry. You can use our Bash 
   script or you can check for more info the 
[official documentation](https://kind.sigs.k8s.io/docs/user/quick-start/#creating-a-cluster):

```bash
$ ./scripts/kind-with-registry.sh
```

> :warning: **Warning 1:** make sure that all pods are up an running before proceeding

## Deploying Dracon dependencies

1. When your Kubernetes cluster is ready to deploy the dependencies, make sure that you have set
   the correct context and then run the following:

```bash
$ make dev-deploy
```

The `dev-deploy` make target will deploy the following components:
a. ArangoDB
b. Nginx Ingress
c. Elastic Search Operator
d. Elastic Search
e. Kibana
f. MongoDB
g. Postgres
h. TektonCD and TektonCD Dashboard

2. Expose the TektonCD Dashboard

```bash
$ kubectl -n tekton-pipelines port-forward svc/tekton-dashboard 9097:9097
```

3. Expose the Kibana Dashboard.

```bash
# Use `kubectl port-forward ...` to access the Kibana UI:
$ kubectl -n dracon port-forward svc/dracon-kb-kibana-kb-http 5601:5601
# You can obtain the password by examining the `dracon-es-elasticsearch-es-elastic-user` secret:
# The username is `elastic`.
$ kubectl -n dracon get secret dracon-es-elasticsearch-es-elastic-user \
          -o=jsonpath='{.data.elastic}' | \
          base64 -d && echo
```

4. Expose the ElasticSearch Dashboard

```bash
# Use `kubectl port-forward ...` to access the Kibana UI:
$ kubectl -n dracon port-forward svc/dracon-kb-kibana-kb-http 5601:5601
```

The username/password is the same as Kibana

## Building all the components

We use [Kustomize Components](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/components.md)
to create composable Dracon Pipelines. These components are packaged into container images using
`make`. The following examples are using the local container registry used by the KiND cluster, but
make sure that you replace the URL with the registry URL that you are using, if you are using
something else:

```bash
$ make publish-component-containers CONTAINER_REPO=localhost:5001/ocurity/dracon
```

## Using a different base image for your images

If you need your images to have a different base image then you can pass the `BASE_IMAGE` variable
to the `components` or `publish-component-containers` to change it to whatever you need. The targets
build the binaries and place them in the `bin` directory and then other targets package them into
containers with `scratch` as the base image.

There are some components that require extra components or special treatment and these components
have their own Makefiles. In those cases you can place a `.custom_image` file in the directory
with the base image you wish to use and that will be picked up by the Makefile and build the
container.

## Running your first Dracon pipeline

1. Create the following simple Dracon Pipeline in your directory:

  ```yaml
  ---
  # ./kustomization.yaml
  apiVersion: kustomize.config.k8s.io/v1beta1
  kind: Kustomization

  nameSuffix: -github-com-kubernetes-kubernetes
  namespace: default

  resources:
  - https://github.com/ocurity/dracon//components/base/

  components:
  - https://github.com/ocurity/dracon/components/sources/git/
  - https://github.com/ocurity/dracon/components/producers/aggregator/
  - https://github.com/ocurity/dracon/components/producers/golang-gosec/
  - https://github.com/ocurity/dracon/components/producers/golang-nancy/
  - https://github.com/ocurity/dracon/components/enrichers/aggregator/
  - https://github.com/ocurity/dracon/components/enrichers/deduplication/
  - https://github.com/ocurity/dracon/components/consumers/elasticsearch/
  ```

2. Run the following to create the Tekton Pipeline, Task, etc. resources on your cluster:

```bash
$ kustomize build | kubectl apply -f -
# Note: you can just run the below to see the generated Tekton Pipeline resources
# $ kustomize build
```

3. Create the following Tekton PipelineRun file:

```yaml
---
# pipelinerun.yaml
# Run `kubectl create ...` with this file.
apiVersion: tekton.dev/v1beta1
kind: PipelineRun
metadata:
  generateName: dracon-github-com-kubernetes-kubernetes-
spec:
  pipelineRef:
    name: dracon-github-com-kubernetes-kubernetes
  params:
    - name: repository_url
      value: https://github.com/kubernetes/kubernetes.git
    - name: consumer-elasticsearch-url
      value: http://quickstart-es-http.default.svc:9200
  workspaces:
    - name: source-code-ws
      subPath: source-code
      volumeClaimTemplate:
        spec:
          accessModes:
            - ReadWriteOnce
          resources:
            requests:
              storage: 1Gi
```

4. Create the PipelineRun resource:

```bash
$ kubectl create -f pipelinerun.yaml
```

5. Observe the PipelineRun at http://localhost:8001/api/v1/namespaces/tekton-pipelines/services/tekton-dashboard:http/proxy/#/about

6. Once the PipelineRun has finished, you can view the output in Kibana at http://localhost:5601.





3. Deploy the Helm chart with the golang project

```bash
$ helm upgrade golang-project-pipeline ./examples/pipelines/golang-project/ \
    --install \
    --namespace dracon
```

4. Start a pipeline run

```bash
$ kubectl apply -n dracon -f ./examples/pipelines/golang-project/pipelinerun/pipelinerun.yaml
```

5. Expose the Tekton dashboard and Kibana UI

```bash
$ kubectl -n tekton-pipelines port-forward svc/tekton-dashboard 9097:9097
$ kubectl -n dracon port-forward svc/dracon-kb-kibana-kb-http 5601:5601
```

6. Get the token to access Kibana

```bash
# The username is `elastic`.
$ kubectl -n dracon get secret dracon-es-elasticsearch-es-elastic-user \
        -o=jsonpath='{.data.elastic}' \
        | base64 -d - \
        | xargs echo "$1"
```

7. Deploy the Golang pipeline



8. Wait for the pipeline to finish running by monitoring it in http://localhost:9097.

9.  Once the pipelinerun has finished running you can view your results in Kibana: http://localhost:5601.