package statusbar

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/commands"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg), nil

	case commands.GetNamespacesMsg:
		return m.handleGetNamespacesMsg(msg), nil
	case commands.SetNamespaceMsg:
		return m.handleSetNamespaceMsg(msg), nil
	case commands.SetKindMsg:
		return m.handleSetKindMsg(msg), nil
	}

	return m, nil

}

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) Model {
	m.SetSize(msg.Width/2, 3)
	return *m
}

func (m *Model) handleGetNamespacesMsg(msg commands.GetNamespacesMsg) Model {
	m.SetNamespace("")
	return *m
}

func (m *Model) handleSetNamespaceMsg(msg commands.SetNamespaceMsg) Model {
	m.SetNamespace(string(msg))
	return *m
}

func (m *Model) handleSetKindMsg(msg commands.SetKindMsg) Model {
	m.SetKind(string(msg))
	return *m
}
