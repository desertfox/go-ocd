package api

import "fmt"

func GetNamespaces() []string {
	var namespaces []string
	for i := 1; i <= 100; i++ {
		namespaces = append(namespaces, fmt.Sprintf("namespace%d", i))
	}
	return namespaces
}

func GetBuildConfig() []string {
	return []string{"bc1", "bc2"}
}
