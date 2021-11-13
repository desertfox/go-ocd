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
	}

	return func() tea.Msg {
		return msgtypes.SetKindMsg{Kind: kind, Items: items}
	}
}

func (m *Model) GetKindInstanceCmd(instance string) tea.Cmd {
	yaml := m.api.GetInstance(m.namespace, m.kind, instance)
	return func() tea.Msg {
		return msgtypes.KindInstanceYaml(yaml)
	}
}
