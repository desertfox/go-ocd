package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-ocd/pkg/list"
)

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.topPane.SetSize(msg.Width, 3)

	m.leftPane.SetSize(msg.Width/2, msg.Height-3)

	m.rightPane.SetSize(msg.Width/2, msg.Height-3)

	if !m.ready {
		m.ready = true
	}

	return m, nil
}

func (m *Model) handleGetNamespacesMsg(msg getNamespacesMsg) (tea.Model, tea.Cmd) {
	m.SetNamespace(namespace(msg))
	m.topPane.SetContent(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Render(string(m.namespace)))

	m.SetKind(kind("namespace"))
	m.list = list.NewModel("namespace")
	m.list.AddItems(m.GetNamespaces())
	m.leftPane.SetContent(m.list.View())

	return m, nil
}

func (m *Model) handleSetNamespaceMsg(msg setNamespaceMsg) (tea.Model, tea.Cmd) {
	m.SetNamespace(namespace(msg))
	m.topPane.SetContent(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Render(string(m.namespace)))

	m.SetKind(kind("kinds"))
	var defaultKinds = []string{"BuildConfig", "ImageStream", "DeploymentConfig"}
	m.list = list.NewModel("kinds")
	m.list.AddItems(defaultKinds)
	m.leftPane.SetContent(m.list.View())

	m.rightPane.SetContent(fmt.Sprintf("Please select %s", m.kind))

	return m, nil
}

func (m *Model) handleSetKindMsg(msg setKindMsg) (tea.Model, tea.Cmd) {
	m.SetKind(kind(msg))
	m.list = list.NewModel(string(m.kind))
	m.list.AddItems(m.GetBuildConfig())
	m.leftPane.SetContent(m.list.View())

	m.rightPane.SetContent(fmt.Sprintf("Please select %s", m.kind))

	return m, nil
}

func (m *Model) handleEnterKey() (tea.Model, tea.Cmd) {
	switch m.kind {
	case "namespace":
		return m, m.setNamespaceCmd(m.list.GetItemAtCursor())
	case "kinds":
		return m, m.setKindCmd(m.list.GetItemAtCursor())
	default:
		m.SetKind(kind(m.list.GetItemAtCursor()))
		m.leftPane.SetContent(m.list.View())
		m.rightPane.SetContent(fmt.Sprintf("describe todo %s", m.kind))
	}
	return m, nil
}

func (m *Model) handleCursorUp() (tea.Model, tea.Cmd) {
	m.list.CursorUp()
	m.leftPane.SetContent(m.list.View())
	return m, nil
}

func (m *Model) handleCursorDown() (tea.Model, tea.Cmd) {
	m.list.CursorDown()
	m.leftPane.SetContent(m.list.View())
	return m, nil
}
