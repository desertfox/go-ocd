package ocd

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/list"
	"github.com/go-ocd/pkg/msgtypes"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case msgtypes.GetNamespacesMsg:
		cmds = append(cmds, m.handleGetNamespacesMsg(msg))

	case msgtypes.SetNamespaceMsg:
		cmds = append(cmds, m.handleSetNamespaceMsg(msg))

	case msgtypes.SetKindMsg:
		m.handleSetKindMsg(msg)

	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			return m.handleEnterKey()

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

func (m *Model) handleGetNamespacesMsg(msg msgtypes.GetNamespacesMsg) tea.Cmd {
	m.SetNamespace("")
	return m.SetKindCmd("namespace")
}

func (m *Model) handleSetNamespaceMsg(msg msgtypes.SetNamespaceMsg) tea.Cmd {
	m.SetNamespace(string(msg))
	return m.SetKindCmd("kind")
}

func (m *Model) handleSetKindMsg(msg msgtypes.SetKindMsg) {
	m.SetKind(string(msg.Kind))

	switch msg.Kind {
	case "namespace":
		m.keys.ShowInstanceKeys(false)
		m.list.SetKind("namespace")
		m.list = list.NewModel(m.kind, m.list.Height, m.list.Width, msg.Items)
	case "kind":
		m.list.SetKind("kind")
		m.list = list.NewModel(m.kind, m.list.Height, m.list.Width, msg.Items)
	//Need to be able to match against all the different kinds
	default: //This is in the context of an instance
		m.keys.ShowInstanceKeys(true)
		m.list.SetKind(string(m.kind))
		m.list = list.NewModel(m.kind, m.list.Height, m.list.Width, msg.Items)
	}

}

func (m *Model) handleEnterKey() (tea.Model, tea.Cmd) {
	switch m.kind {
	case "namespace":
		return m, m.SetNamespaceCmd(m.list.GetItemAtCursor())
	case "kind":
		return m, m.SetKindCmd(m.list.GetItemAtCursor())
	default:
		m.keys.ShowSelectKeys(false)
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
		m.keys.ShowInstanceKeys(false)
		m.keys.ShowSelectKeys(true)
		return m, m.SetKindCmd("kind")
	}
}
