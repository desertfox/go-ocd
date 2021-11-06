package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/api"
)

func SetNamespaceCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return SetNamespaceMsg(namespace)
	}
}

func GetNamespacesCmd() tea.Cmd {
	fakeNamespaces := api.GetNamespaces()

	return func() tea.Msg {
		return GetNamespacesMsg(fakeNamespaces)
	}
}

func SetKindCmd(kind string) tea.Cmd {
	return func() tea.Msg {
		return SetKindMsg(kind)
	}
}

func GetKindInstanceDescribeCmd(instance string) tea.Cmd {
	return func() tea.Msg {
		return GetKindInstanceDescribeMsg("describe")
	}
}
