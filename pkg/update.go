package ocd

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/msgtypes"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case msgtypes.GetNamespacesMsg:
		m.handleGetNamespacesMsg(&cmds, msg)

	case msgtypes.SetNamespaceMsg:
		m.handleSetNamespaceMsg(&cmds, msg)

	case msgtypes.SetKindMsg:
		m.handleSetKindMsg(msg)

	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Enter):
			return m.handleEnterKey()

		case key.Matches(msg, m.keys.Delete):
			return m, nil

		case key.Matches(msg, m.keys.Back):
			return m.handleGoBack()
		}
	}

	m.statusbar, cmd = m.statusbar.Update(msg)
	cmds = append(cmds, cmd)

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	m.pane, cmd = m.pane.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) {
	if !m.ready {
		m.ready = true
	}
}

func (m *Model) handleGetNamespacesMsg(cmds *[]tea.Cmd, msg msgtypes.GetNamespacesMsg) {
	m.SetNamespace("")
	*cmds = append(*cmds, m.SetKindCmd("namespace"))
}

func (m *Model) handleSetNamespaceMsg(cmds *[]tea.Cmd, msg msgtypes.SetNamespaceMsg) {
	m.SetNamespace(string(msg))
	*cmds = append(*cmds, m.SetKindCmd("kind"))
}

func (m *Model) handleSetKindMsg(msg msgtypes.SetKindMsg) {
	m.SetKind(string(msg.Kind))

	switch msg.Kind {
	case "namespace":
		m.list.SetKind("namespace")
		m.list.AddItems(msg.Items)
	case "kind":
		m.list.SetKind("kind")
		m.list.AddItems(msg.Items)
	default:
		m.list.SetKind(string(m.kind))
		m.list.AddItems(msg.Items)
	}

}

func (m *Model) handleEnterKey() (tea.Model, tea.Cmd) {
	switch m.kind {
	case "namespace":
		return m, m.SetNamespaceCmd(m.list.GetItemAtCursor())
	case "kind":
		return m, m.SetKindCmd(m.list.GetItemAtCursor())
	default:
		return m, m.GetKindInstanceCmd(m.list.GetItemAtCursor())
	}
}

func (m *Model) handleGoBack() (tea.Model, tea.Cmd) {
	switch m.kind {
	case "namespace":
		return m, m.GetNamespacesCmd()
	case "kind":
		return m, m.GetNamespacesCmd()
	default:
		return m, m.SetKindCmd("kind")
	}
}
