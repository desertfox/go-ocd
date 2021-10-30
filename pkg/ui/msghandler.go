package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/list"
)

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.topPane.SetSize(msg.Width, 3)

	m.leftPane.SetSize(msg.Width/2, msg.Height-3)

	m.rightPane.SetSize(msg.Width/2, msg.Height-3)

	m.help.Width = msg.Width/2 - 10

	m.buildRightPane()

	if !m.ready {
		m.ready = true
	}

	return m, nil
}

func (m *Model) handleGetNamespacesMsg(msg getNamespacesMsg) (tea.Model, tea.Cmd) {
	m.SetNamespace(namespace(msg))
	m.SetKind(kind("namespace"))
	m.list = list.NewModel("namespace")
	m.list.AddItems(m.GetNamespaces())
	return m, m.batchAllPanes()
}

func (m *Model) handleSetNamespaceMsg(msg setNamespaceMsg) (tea.Model, tea.Cmd) {
	m.SetNamespace(namespace(msg))
	m.SetKind(kind("kinds"))
	m.list = list.NewModel("kinds")
	m.list.AddItems([]string{"BuildConfig", "ImageStream", "DeploymentConfig"})
	return m, m.batchAllPanes()
}

func (m *Model) handleSetKindMsg(msg setKindMsg) (tea.Model, tea.Cmd) {
	m.SetKind(kind(msg))
	m.list = list.NewModel(string(m.kind))
	m.list.AddItems(m.GetBuildConfig())
	return m, m.batchAllPanes()
}

func (m *Model) handleEnterKey() (tea.Model, tea.Cmd) {
	switch m.kind {
	case "namespace":
		return m, m.setNamespaceCmd(m.list.GetItemAtCursor())
	case "kinds":
		return m, m.setKindCmd(m.list.GetItemAtCursor())
	default:
		m.SetKind(kind(m.list.GetItemAtCursor()))
		return m, m.batchAllPanes()
	}
}

func (m *Model) handleCursorUp() (tea.Model, tea.Cmd) {
	m.list.CursorUp()
	return m, m.buildPaneCmd("left")
}

func (m *Model) handleCursorDown() (tea.Model, tea.Cmd) {
	m.list.CursorDown()
	return m, m.buildPaneCmd("left")
}
