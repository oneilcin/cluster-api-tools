# cluster-api-tools

## genClusterApiServerYaml

This tool will use the cluster-api server deploy [template](https://github.com/kubernetes-sigs/cluster-api/blob/master/clusterctl/clusterdeployer/clusterapiservertemplate.go) defined in 
sigs.k8s.io/cluster-api/clusterctl/clusterdeployer, 
create the needed certs for the cluster-api server,
and output the completed template needed to deploy the cluster-api server.

## Run the tool to generate the clusterapi-apiserver.yaml 
    cd $GOPATH/src/sigs.k8s.io
    git clone https://github.com/oneilcin/cluster-api-tools
    cd cluster-api-tools
    go run genClusterApiServerYaml.go > clusterapi-apiserver.yaml  (use any output file name)

## Usage

First deploy the cluster-api server, and then deploy the chosen provider components.

    minikube start --bootstrapper=kubeadm
    kubectl create -f apiserverdeploy.yaml
    kubectl create -f your-provider-components.yaml   (this includes the provider machine and cluster controllers)

Note: See this example for a [provider-components-template](https://github.com/samsung-cnct/cluster-api-provider-ssh/blob/master/clusterctl/examples/ssh/provider-components.yaml.template)
