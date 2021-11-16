package pane

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/helpers"
	"github.com/go-ocd/pkg/msgtypes"
)

func (m Model) Update(msg tea.Msg, disableKeys bool) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg)
	case msgtypes.SetKindMsg:
		return m.handleSetKindMsg(msg)
	case msgtypes.GetKindInstanceDescribeMsg:
		return m.handleGetKindInstanceDescribeMsg(msg)
	case msgtypes.KindInstanceYaml:
		return m.handleKindInstanceYamlMsg(msg)
	}

	if !m.selected {
		return m, nil
	}

	return m, m.updateViewport(msg)
}

func (m *Model) updateViewport(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return cmd
}

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) {
	m.SetSize(helpers.NewDimensions(msg.Width/2-10, msg.Height-9))
}

func (m *Model) handleSetKindMsg(msg msgtypes.SetKindMsg) (Model, tea.Cmd) {
	switch msg.Kind {
	case "namespace":

	case "kind":
		m.SetContent(fmt.Sprintf("Please select a %s", msg.Kind))
	default:
		m.SetContent(fmt.Sprintf("Please select a %s instance", msg.Kind))
	}

	return *m, nil
}

func (m *Model) handleGetKindInstanceDescribeMsg(describe msgtypes.GetKindInstanceDescribeMsg) (Model, tea.Cmd) {
	m.viewport.SetContent(string(describe))
	return *m, nil
}

func (m *Model) handleKindInstanceYamlMsg(yaml msgtypes.KindInstanceYaml) (Model, tea.Cmd) {
	m.viewport.SetContent(string(yaml))
	return *m, nil
}
