package ocd

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/msgtypes"
)

func (m *Model) SetNamespaceCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return msgtypes.SetNamespaceMsg(namespace)
	}
}

func (m *Model) GetNamespacesCmd() tea.Cmd {
	namespaces := m.api.GetNamespaces()

	return func() tea.Msg {
		return msgtypes.GetNamespacesMsg(namespaces)
	}
}

func (m *Model) SetKindCmd(kind string) tea.Cmd {
	var items []string

	switch strings.ToLower(kind) {
	case "namespace":
		items = m.api.GetNamespaces()
	case "kind":
		items = m.api.GetAvailKinds()
	case "buildconfig":
		items = m.api.GetBuildConfigs()
	case "secrets":
		items = m.api.GetSecrets()
	default:
		items = m.api.GetBuildConfig()
	}

	return func() tea.Msg {
		return msgtypes.SetKindMsg{Kind: kind, Items: items}
	}
}

func (m *Model) GetKindInstanceCmd(instance string) tea.Cmd {
	//sting api.GetKind(m.Kind, instance)
	return func() tea.Msg {
		return msgtypes.KindInstanceYaml("lots of yaml here")
	}
}

func (m *Model) GetKindInstanceDescribeCmd(instance string) tea.Cmd {
	return func() tea.Msg {
		return msgtypes.GetKindInstanceDescribeMsg(instance)
	}
}
