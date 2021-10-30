package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case getNamespacesMsg:
		return m.handleGetNamespacesMsg(msg)

	case setNamespaceMsg:
		return m.handleSetNamespaceMsg(msg)

	case setKindMsg:
		return m.handleSetKindMsg(msg)

	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg)

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = true

		case key.Matches(msg, m.keys.Up):
			return m.handleCursorUp()

		case key.Matches(msg, m.keys.Down):
			return m.handleCursorDown()

		case key.Matches(msg, m.keys.Enter):
			return m.handleEnterKey()

		case key.Matches(msg, m.keys.Delete):
			return m, nil
		}

	}

	return m, tea.Batch(cmds...)
}
