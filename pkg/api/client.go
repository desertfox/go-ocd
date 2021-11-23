package api

import (
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
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
	return []string{"BuildConfig", "ImageStream", "DeploymentConfig", "Route", "Secret", "Service", "Endpoint", "ResourceQuota", "Node", "Pod", "ServiceAccount"}
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

func (client Client) GetInstances(namespace, kind string) []string {
	switch strings.ToLower(kind) {
	case "buildconfig":
		return client.GetBuildConfigs(namespace)

	case "secret":
		return client.GetSecrets(namespace)
	}

	return []string{""}
}

func (client Client) parseResourceItems(item interface{}) []string {
	var list []string

	switch item := item.(type) {
	case v1.SecretList:
		for _, i := range item.Items {
			list = append(list, i.Name)
		}
	case v1.ResourceQuotaList:
		for _, i := range item.Items {
			list = append(list, i.Name)
		}
	case v1.PodList:
		for _, i := range item.Items {
			list = append(list, i.Name)
		}
	case v1.EndpointsList:
		for _, i := range item.Items {
			list = append(list, i.Name)
		}
	case v1.NodeList:
		for _, i := range item.Items {
			list = append(list, i.Name)
		}
	case v1.ConfigMapList:
		for _, i := range item.Items {
			list = append(list, i.Name)
		}
	case v1.PersistentVolumeList:
		for _, i := range item.Items {
			list = append(list, i.Name)
		}
	}

	return list
}
