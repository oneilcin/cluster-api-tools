# cluster-api-tools

## genClusterApiServerYaml

This tool will use the cluster-api/pkg/deployer code to generate the cluster-api apiserver deployment manifest.

## Run the tool to generate the clusterapi-apiserver.yaml

    cd $GOPATH/src/sigs.k8s.io
    git clone https://github.com/oneilcin/cluster-api-tools
    cd cluster-api-tools
    go run genClusterApiServerYaml.go > clusterapi-apiserver.yaml

## Usage

First deploy the cluster-api server, and then deploy the chosen provider components.

    minikube start --bootstrapper=kubeadm
    kubectl create -f clusterapi-apiserver.yaml -f provider-components.yaml
    # wait for the apiserver pod to be ready
    kubectl get pods -w
    kubectl create -f cluster.yaml -f machines.yaml --validate=false

Instructions to generate the provider-components.yaml are [here](https://github.com/samsung-cnct/cluster-api-provider-ssh/blob/master/clusterctl/examples/ssh/README.md)

Note: See this example for a [provider-components-template](https://github.com/samsung-cnct/cluster-api-provider-ssh/blob/master/clusterctl/examples/ssh/provider-components.yaml.template)
