package api

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func (client Client) GetSecrets(namespace string) []string {
	if client.fake {
		return makeFakeList("secrets")
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	secrets, _ := kclient.Secrets(namespace).List(context.TODO(), metav1.ListOptions{})

	var secretList []string
	for _, s := range secrets.Items {
		secretList = append(secretList, s.Name)
	}

	return secretList
}

func (client Client) GetSecret(namespace, name string) string {
	if client.fake {
		return "secretstuff"
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	secret, _ := kclient.Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return fmt.Sprintf("%v", secret)
}
