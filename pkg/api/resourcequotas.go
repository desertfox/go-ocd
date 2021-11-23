package api

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func (client Client) ResourceQuotass(namespace string) []string {
	if client.fake {
		return makeFakeList("PersistentVolumeClaims")
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	rq, _ := kclient.PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})

	var rqList []string
	for _, s := range rq.Items {
		rqList = append(rqList, s.Name)
	}

	return rqList
}

func (client Client) ResourceQuotas(namespace, name string) string {
	if client.fake {
		return "PersistentVolumeClaims"
	}

	kclient, _ := v1.NewForConfig(client.k8client)
	rq, _ := kclient.Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return fmt.Sprintf("%v", rq)
}
