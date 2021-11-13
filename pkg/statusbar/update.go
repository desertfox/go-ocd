package statusbar

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/msgtypes"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg)

	case msgtypes.GetNamespacesMsg:
		return m.handleGetNamespacesMsg(msg)

	case msgtypes.SetNamespaceMsg:
		return m.handleSetNamespaceMsg(msg)

	case msgtypes.SetKindMsg:
		return m.handleSetKindMsg(msg)
	}

	return m, nil

}

func (m Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (Model, tea.Cmd) {
	m.SetSize(msg.Width/2, 3)
	return m, nil
}

func (m Model) handleGetNamespacesMsg(msg msgtypes.GetNamespacesMsg) (Model, tea.Cmd) {
	m.SetNamespace("")
	return m, nil
}

func (m Model) handleSetNamespaceMsg(msg msgtypes.SetNamespaceMsg) (Model, tea.Cmd) {
	m.SetNamespace(string(msg))
	return m, nil
}

func (m Model) handleSetKindMsg(msg msgtypes.SetKindMsg) (Model, tea.Cmd) {
	m.SetKind(string(msg.Kind))
	return m, nil
}
