package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case selectNamespaceKindMsg:
		m.lookupNamespaceKind(string(msg))
		m.leftPane.SetContent(m.list.View())
	case selectNamespaceMsg:
		m.lookupNamespace()
		m.leftPane.SetContent(m.list.View())

	case tea.WindowSizeMsg:
		m.leftPane.SetSize(msg.Width/2, msg.Height)
		m.rightPane.SetSize(msg.Width/2, msg.Height)

		//m.leftPane.SetContent(m.namespace)
		//m.rightPane.SetContent("")

		if !m.ready {
			m.ready = true
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "exit"))):
			return m, tea.Quit
		}

	}

	return m, nil
}
