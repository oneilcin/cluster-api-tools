# cluster-api-tools

## genClusterApiServerYaml

This tool will use the cluster-api/pkg/deployer code to generate the cluster-api apiserver deployment manifest.

## Run the tool to generate the clusterapi-apiserver.yaml

    mkdir -p $GOPATH/src/github.com/oneilcin
    cd $GOPATH/src/github.com/oneilcin
    git clone https://github.com/oneilcin/cluster-api-tools
    cd cluster-api-tools
    go run genClusterApiServerYaml.go > clusterapi-apiserver.yaml

## Usage

### Deploy into a Manager Cluster
The manager cluster can be any kubeadm created kubernetes cluster.

#### Using minikube as your Manager Cluster

Note: There is a [known issue](https://github.com/kubernetes/minikube/issues/3076) with minikube versions later than 0.28.0 not working with the cluster-api apiserver.  Be sure and use minikube version 0.28.0.

     minikube start --bootstrapper=kubeadm

Instructions to generate the provider-components.yaml for the [cluster-api-provider-ssh](https://github.com/samsung-cnct/cluster-api-provider-ssh/blob/master/clusterctl/examples/ssh/README.md)

First deploy the cluster-api server, and then deploy the chosen provider components.
    
    kubectl create -f clusterapi-apiserver.yaml -f provider-components.yaml
    # wait for the apiserver pod to be ready
    kubectl get pods -w


### Create new Managed Clusters using the cluster-api-provider-ssh API (direct, or kubectl)

It is recommended to create each cluster in a new namespace.  Sample manifests for [create cluster](https://github.com/samsung-cnct/cluster-api-provider-ssh/tree/master/assets)

    # create a namespace for each new cluster you are creating
    kubectl create namespace test1
 
    # make sure the manifests below are in the same namespace
    kubectl create -f cluster-private-key.yaml
    kubectl create -f cluster.yaml --validate=false
    kubectl create -f machines-formatted.yaml --validate=false
