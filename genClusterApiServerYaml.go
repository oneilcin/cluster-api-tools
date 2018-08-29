package main

import (
	"fmt"

	"github.com/golang/glog"
	"sigs.k8s.io/cluster-api/pkg/deployer"
)

func main() {
	yaml, err := deployer.GetApiServerYaml()
	if err != nil {
		glog.Error(err)
	}
	fmt.Println(yaml)
}
