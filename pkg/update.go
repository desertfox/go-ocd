package ocd

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/commands"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	log.Printf("%t%#v", msg, msg)

	switch msg := msg.(type) {

	case commands.GetNamespacesMsg:
		m.handleGetNamespacesMsg(&cmds, msg)

	case commands.SetNamespaceMsg:
		m.handleSetNamespaceMsg(&cmds, msg)

	case commands.SetKindMsg:
		m.handleSetKindMsg(msg)

	case commands.GetKindInstanceDescribeMsg:
		m.handleGetKindInstanceDescribeMsg(msg)

	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

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

func (m *Model) handleGetNamespacesMsg(cmds *[]tea.Cmd, msg commands.GetNamespacesMsg) {
	m.SetNamespace("")
	*cmds = append(*cmds, commands.SetKindCmd("namespace"))
}

func (m *Model) handleSetNamespaceMsg(cmds *[]tea.Cmd, msg commands.SetNamespaceMsg) {
	m.SetNamespace(string(msg))
	*cmds = append(*cmds, commands.SetKindCmd("kind"))
}

func (m *Model) handleSetKindMsg(msg commands.SetKindMsg) {
	m.SetKind(string(msg))
}

func (m *Model) handleEnterKey() (tea.Model, tea.Cmd) {
	switch m.kind {
	case "namespace":
		return m, commands.SetNamespaceCmd(m.list.GetItemAtCursor())
	case "kind":
		return m, commands.SetKindCmd(m.list.GetItemAtCursor())
	default:
		return m, commands.GetKindInstanceDescribeCmd(m.list.GetItemAtCursor())
	}
}

func (m *Model) handleGetKindInstanceDescribeMsg(describe commands.GetKindInstanceDescribeMsg) (tea.Model, tea.Cmd) {
	m.pane.SetContent(string(describe))
	return m, nil
}

func (m *Model) handleGoBack() (tea.Model, tea.Cmd) {
	switch m.kind {
	case "namespace":
		return m, commands.GetNamespacesCmd()
	case "kind":
		return m, commands.GetNamespacesCmd()
	default:
		return m, commands.SetKindCmd("kind")
	}
}
