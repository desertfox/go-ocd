package api

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	k8client *rest.Config
}

func NewClient(kubeconfig string) Client {
	k8client, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)

	return Client{k8client: k8client}
}
