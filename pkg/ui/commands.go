package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type setNamespaceMsg string
type getNamespacesMsg string
type setKindMsg string

func (m Model) setNamespaceCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return setNamespaceMsg(namespace)
	}
}

func (m Model) getNamespacesCmd() tea.Cmd {
	return func() tea.Msg {
		return getNamespacesMsg("")
	}
}

func (m Model) setKindCmd(kind string) tea.Cmd {
	return func() tea.Msg {
		return setKindMsg(kind)
	}
}
