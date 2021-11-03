package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case buildPaneMsg:
		return m.buildPane(msg)

	case getNamespacesMsg:
		return m.handleGetNamespacesMsg(msg)

	case setNamespaceMsg:
		return m.handleSetNamespaceMsg(msg)

	case setKindMsg:
		return m.handleSetKindMsg(msg)

	case getKindInstanceDescribeMsg:
		return m.handleGetKindInstanceDescribeMsg(msg)

	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg)

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, m.buildPaneCmd("help")

		case key.Matches(msg, m.keys.Up):
			_, cmd = m.handleCursorUp()
			cmds = append(cmds, cmd)
			//return m.handleCursorUp()

		case key.Matches(msg, m.keys.Down):
			_, cmd = m.handleCursorDown()
			cmds = append(cmds, cmd)
			//return m.handleCursorDown()

		case key.Matches(msg, m.keys.Enter):
			return m.handleEnterKey()
		case key.Matches(msg, m.keys.Delete):
			return m, nil
		case key.Matches(msg, m.keys.Back):
			return m.handleGoBack()
		}

	}

	cmds = append(cmds, m.UpdatePanes(msg))

	return m, tea.Batch(cmds...)
}
