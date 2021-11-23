package api

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func (client Client) GetConfigMaps(namespace string) []string {
	if client.fake {
		return makeFakeList("configmap")
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	configmaps, _ := kclient.ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})

	return client.parseResourceItems(configmaps)
}

func (client Client) GetConfigMap(namespace, name string) string {
	if client.fake {
		return "configmap"
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	configmap, _ := kclient.ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return fmt.Sprintf("%v", configmap)
}
