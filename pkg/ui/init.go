package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	namespace := ""

	if len(os.Args) > 1 {
		namespace = os.Args[1]
	}

	switch {
	case namespace != "":
		cmds = append(cmds, m.setNamespaceCmd(namespace))
	default:
		cmds = append(cmds, m.getNamespacesCmd())
	}

	return tea.Batch(cmds...)
}
