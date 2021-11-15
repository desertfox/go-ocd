package api

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	buildv1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
)

func (client Client) GetBuildConfigs(namespace string) []string {

	if client.fake {
		return makeFakeList("buildconfig")
	}

	buildV1Client, _ := buildv1.NewForConfig(client.k8client)
	builds, _ := buildV1Client.Builds(namespace).List(context.TODO(), metav1.ListOptions{})

	var buildconfigs []string = []string{}

	for _, build := range builds.Items {
		buildconfigs = append(buildconfigs, build.Name)
	}

	return buildconfigs
}

func (client Client) GetBuildConfigInstance(namespace, instanceName string) string {
	//return fmt.Sprintf("%s/BuildConfig/%s", namespace, instanceName)
	return fakeYaml
}
