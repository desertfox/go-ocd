package statusbar

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/commands"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg)

	case commands.GetNamespacesMsg:
		return m.handleGetNamespacesMsg(msg)

	case commands.SetNamespaceMsg:
		return m.handleSetNamespaceMsg(msg)

	case commands.SetKindMsg:
		return m.handleSetKindMsg(msg)
	}

	return m, nil

}

func (m Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (Model, tea.Cmd) {
	m.SetSize(msg.Width/2, 3)
	return m, nil
}

func (m Model) handleGetNamespacesMsg(msg commands.GetNamespacesMsg) (Model, tea.Cmd) {
	m.SetNamespace("")
	return m, nil
}

func (m Model) handleSetNamespaceMsg(msg commands.SetNamespaceMsg) (Model, tea.Cmd) {
	m.SetNamespace(string(msg))
	return m, nil
}

func (m Model) handleSetKindMsg(msg commands.SetKindMsg) (Model, tea.Cmd) {
	m.SetKind(string(msg))
	return m, nil
}
