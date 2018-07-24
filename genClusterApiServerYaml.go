package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"text/template"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/cert"
	"k8s.io/client-go/util/cert/triple"
	"sigs.k8s.io/cluster-api/clusterctl/clusterdeployer"
)

var apiServerImage = "gcr.io/k8s-cluster-api/cluster-apiserver:0.0.5"

type caCertParams struct {
	caBundle string
	tlsCrt   string
	tlsKey   string
}

func getApiServerCerts() (*caCertParams, error) {
	const name = "clusterapi"
	const namespace = corev1.NamespaceDefault

	caKeyPair, err := triple.NewCA(fmt.Sprintf("%s-certificate-authority", name))
	if err != nil {
		return nil, fmt.Errorf("failed to create root-ca: %v", err)
	}

	apiServerKeyPair, err := triple.NewServerKeyPair(
		caKeyPair,
		fmt.Sprintf("%s.%s.svc", name, namespace),
		name,
		namespace,
		"cluster.local",
		[]string{},
		[]string{})
	if err != nil {
		return nil, fmt.Errorf("failed to create apiserver key pair: %v", err)
	}

	certParams := &caCertParams{
		caBundle: base64.StdEncoding.EncodeToString(cert.EncodeCertPEM(caKeyPair.Cert)),
		tlsKey:   base64.StdEncoding.EncodeToString(cert.EncodePrivateKeyPEM(apiServerKeyPair.Key)),
		tlsCrt:   base64.StdEncoding.EncodeToString(cert.EncodeCertPEM(apiServerKeyPair.Cert)),
	}
	return certParams, nil
}

func getApiServerYaml() (string, error) {
	tmpl, err := template.New("config").Parse(clusterdeployer.ClusterAPIAPIServerConfigTemplate)
	if err != nil {
		return "", err
	}

	certParms, err := getApiServerCerts()
	if err != nil {
		glog.Errorf("Error: %v", err)
		return "", err
	}

	type params struct {
		APIServerImage string
		CABundle       string
		TLSCrt         string
		TLSKey         string
	}

	var tmplBuf bytes.Buffer
	err = tmpl.Execute(&tmplBuf, params{
		APIServerImage: apiServerImage,
		CABundle:       certParms.caBundle,
		TLSCrt:         certParms.tlsCrt,
		TLSKey:         certParms.tlsKey,
	})
	if err != nil {
		return "", err
	}

	return string(tmplBuf.Bytes()), nil
}

func main() {
	yaml, err := getApiServerYaml()
	if err != nil {
		glog.Error(err)
	}
	fmt.Println(yaml)
}
