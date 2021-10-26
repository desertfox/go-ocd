package ui

import tea "github.com/charmbracelet/bubbletea"

type selectNamespaceMsg string

type listNamespaceMsg bool

func (m Model) selectNamespaceCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return selectNamespaceMsg(namespace)
	}
}

func (m Model) listNamespaceCmd() tea.Cmd {
	return func() tea.Msg {
		return listNamespaceMsg(true)
	}
}
