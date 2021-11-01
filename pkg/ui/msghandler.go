package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/list"
)

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.panes["selected"].SetSize(msg.Width/2, 2)

	m.panes["list"].SetSize(msg.Width/2, msg.Height-m.panes["selected"].Height()-32)

	m.panes["help"].SetSize(msg.Width/2, 30)

	m.panes["preview"].SetSize(msg.Width/2, msg.Height-m.panes["selected"].Height())

	if !m.ready {
		m.ready = true
	}

	return m, m.batchAllPanes()
}

/*
- namespace
- kind
- $kind
- instance-of-$kind
*/
func (m *Model) handleGetNamespacesMsg(msg getNamespacesMsg) (tea.Model, tea.Cmd) {
	m.namespace = namespace("")

	return m, tea.Sequentially(
		m.setKindCmd("namespace"),
		m.batchAllPanes(),
	)
}

func (m *Model) handleSetNamespaceMsg(msg setNamespaceMsg) (tea.Model, tea.Cmd) {
	m.namespace = namespace(msg)

	return m, tea.Sequentially(
		m.setKindCmd("kind"),
		m.batchAllPanes(),
	)
}

func (m *Model) handleSetKindMsg(msg setKindMsg) (tea.Model, tea.Cmd) {
	m.kind = kind(msg)

	switch msg {
	case "namespace":
		m.list = list.NewModel("namespace")
		m.list.AddItems(m.GetNamespaces())
	case "kind":
		m.list = list.NewModel("kind")
		m.list.AddItems([]string{"BuildConfig", "ImageStream", "DeploymentConfig"})
		m.preview = fmt.Sprintf("Please select a %s", m.kind)
	default:
		m.list = list.NewModel(string(m.kind))
		m.list.AddItems(m.GetBuildConfig())
		m.preview = fmt.Sprintf("Please select a %s instance", m.kind)
	}

	return m, m.batchAllPanes()
}

func (m *Model) handleEnterKey() (tea.Model, tea.Cmd) {
	switch m.kind {
	case "namespace":
		return m, m.setNamespaceCmd(m.list.GetItemAtCursor())
	case "kind":
		return m, m.setKindCmd(m.list.GetItemAtCursor())
	default:
		return m, m.getKindInstanceDescribeCmd(m.list.GetItemAtCursor())
	}
}

func (m *Model) handleCursorUp() (tea.Model, tea.Cmd) {
	m.list.CursorUp()
	return m, m.batchAllPanes()
}

func (m *Model) handleCursorDown() (tea.Model, tea.Cmd) {
	m.list.CursorDown()
	return m, m.batchAllPanes()
}

func (m *Model) handleGetKindInstanceDescribeMsg(describe getKindInstanceDescribeMsg) (tea.Model, tea.Cmd) {
	m.preview = string(describe)
	return m, m.batchAllPanes()
}
