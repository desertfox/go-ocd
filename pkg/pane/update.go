package pane

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/commands"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg)
	case commands.SetKindMsg:
		return m.handleSetKindMsg(msg)
	}

	return m, m.updateViewport(msg)
}

func (m *Model) updateViewport(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return cmd
}

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) {
	m.SetSize(msg.Width/2-10, msg.Height-9)
}

func (m *Model) handleSetKindMsg(msg commands.SetKindMsg) (Model, tea.Cmd) {
	switch msg {
	case "namespace":

	case "kind":
		m.SetContent(fmt.Sprintf("Please select a %s", msg))
	default:
		m.SetContent(fmt.Sprintf("Please select a %s instance", msg))
	}

	return *m, nil
}
