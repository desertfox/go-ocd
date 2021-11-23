package ocd

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	switch {
	case m.namespace != "":
		cmds = append(cmds, m.SetNamespaceCmd(m.namespace))
	default:
		cmds = append(cmds, m.GetNamespacesCmd())
	}

	return tea.Batch(cmds...)
}
