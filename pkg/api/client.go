package api

import (
	"fmt"
	"strings"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	k8client *rest.Config
	fake     bool
}

func NewClient(kubeconfig string, fake bool) Client {
	k8client, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)

	return Client{k8client, fake}
}

func (client Client) GetAvailKinds() []string {
	return []string{"BuildConfig", "ImageStream", "DeploymentConfig", "Routes", "Secrets", "Services"}
}

func (client Client) GetInstance(namespace, kind, instance string) string {

	switch strings.ToLower(kind) {
	case "buildconfig":
		return client.GetBuildConfig(namespace, instance)

	case "secret":
		return client.GetSecret(namespace, instance)
	}

	return fmt.Sprintf("No supporting kind found for %s in namespace %s", kind, namespace)
}
