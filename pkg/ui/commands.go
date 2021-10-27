package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type setNamespaceMsg string
type setDefaultKindMsg bool

func (m Model) setNamespaceCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return setNamespaceMsg(namespace)
	}
}

func (m Model) setDefaultKindCmd() tea.Cmd {
	return func() tea.Msg {
		return setDefaultKindMsg(true)
	}
}
