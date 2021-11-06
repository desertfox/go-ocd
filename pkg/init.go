package ocd

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/commands"
)

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	namespace := ""

	if len(os.Args) > 1 {
		namespace = os.Args[1]
	}

	switch {
	case namespace != "":
		cmds = append(cmds, commands.SetNamespaceCmd(namespace))
	default:
		cmds = append(cmds, commands.GetNamespacesCmd())
	}

	return tea.Batch(cmds...)
}
