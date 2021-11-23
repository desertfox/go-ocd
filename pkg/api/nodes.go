package api

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func (client Client) GetNodes(namespace string) []string {
	if client.fake {
		return makeFakeList("Nodes")
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	nodes, _ := kclient.Nodes().List(context.TODO(), metav1.ListOptions{})

	var nodesList []string
	for _, s := range nodes.Items {
		nodesList = append(nodesList, s.Name)
	}

	return nodesList
}

func (client Client) GetNode(namespace, name string) string {
	if client.fake {
		return "Nodestuff"
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	node, _ := kclient.Nodes().Get(context.TODO(), name, metav1.GetOptions{})

	return fmt.Sprintf("%v", node)
}
