package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case setDefaultKindMsg:
		m.list.SetDefaultKind()
		m.leftPane.SetContent(m.list.View())

	case setNamespaceMsg:
		m.list.SetNamespace(string(msg))
		m.leftPane.SetContent(m.list.View())

	case tea.WindowSizeMsg:
		m.leftPane.SetSize(msg.Width/2, msg.Height)
		m.rightPane.SetSize(msg.Width/2, msg.Height)

		if !m.ready {
			m.ready = true
		}
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "exit"))):
			return m, tea.Quit

		case key.Matches(msg, key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Scroll up"))):
			m.list.CursorUp()
			m.leftPane.SetContent(m.list.View())

		case key.Matches(msg, key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Scroll down"))):
			m.list.CursorDown()
			m.leftPane.SetContent(m.list.View())

		}
	}

	return m, nil
}
