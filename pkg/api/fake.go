package api

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed fakebuildconfig.yml
	fakeYaml string
)

func makeFakeList(str string) []string {
	var kind []string
	for i := 1; i <= 100; i++ {
		kind = append(kind, fmt.Sprintf("%s-%d", str, i))
	}
	return kind
}
