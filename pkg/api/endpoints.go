package api

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func (client Client) GetEndpointss(namespace string) []string {
	if client.fake {
		return makeFakeList("Endpoints")
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	endpoints, _ := kclient.Endpoints(namespace).List(context.TODO(), metav1.ListOptions{})

	var EndpointsList []string
	for _, s := range endpoints.Items {
		EndpointsList = append(EndpointsList, s.Name)
	}

	return EndpointsList
}

func (client Client) GetEndpoints(namespace, name string) string {
	if client.fake {
		return "Endpoints"
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	endpoints, _ := kclient.Endpoints(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return fmt.Sprintf("%v", endpoints)
}
