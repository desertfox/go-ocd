package ocd

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/msgtypes"
)

func SetNamespaceCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return msgtypes.SetNamespaceMsg(namespace)
	}
}

type Apier interface {
	GetNamespaces() []string
	GetAvailKinds() []string
	GetBuildConfigs() []string
	GetBuildConfig() []string
	GetSecrets() []string
}

func GetNamespacesCmd(api Apier) tea.Cmd {
	fakeNamespaces := api.GetNamespaces()

	return func() tea.Msg {
		return msgtypes.GetNamespacesMsg(fakeNamespaces)
	}
}

func SetKindCmd(kind string, api Apier) tea.Cmd {
	var items []string

	switch strings.ToLower(kind) {
	case "namespace":
		items = api.GetNamespaces()
	case "kind":
		items = api.GetAvailKinds()
	case "buildconfig":
		items = api.GetBuildConfigs()
	case "secrets":
		items = api.GetSecrets()
	default:
		items = api.GetBuildConfig()
	}

	return func() tea.Msg {
		return msgtypes.SetKindMsg{Kind: kind, Items: items}
	}
}

func GetKindInstanceCmd(instance string) tea.Cmd {
	//sting api.GetKind(m.Kind, instance)
	return func() tea.Msg {
		return msgtypes.KindInstanceYaml("lots of yaml here")
	}
}

func GetKindInstanceDescribeCmd(instance string) tea.Cmd {
	return func() tea.Msg {
		return msgtypes.GetKindInstanceDescribeMsg(instance)
	}
}
