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

		case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "exit"))):
			return m, tea.Quit

		case key.Matches(msg, key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Scroll up"))):
			return m.handleCursorUp()

		case key.Matches(msg, key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Scroll down"))):
			return m.handleCursorDown()

		case key.Matches(msg, key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "Select"))):
			return m.handleEnterKey()
		}
	}

	return m, tea.Batch(cmds...)
}
