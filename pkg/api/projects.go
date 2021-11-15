package api

import (
	"context"

	projectv1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client Client) GetProjects() []string {

	if client.fake {
		return makeFakeList("namespaces")
	}

	var namespaces []string = []string{}

	projectV1Client, _ := projectv1.NewForConfig(client.k8client)
	projects, _ := projectV1Client.Projects().List(context.TODO(), metav1.ListOptions{})

	for _, project := range projects.Items {
		namespaces = append(namespaces, project.Name)
	}
	return namespaces
}
