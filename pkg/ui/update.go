package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case buildPaneMsg:
		return m.BuildPane(msg)

	case getNamespacesMsg:
		return m.BuildNamespacesList()

	case setNamespaceMsg:
		return m.SetNamespaceAndBuildKindsList(msg)

	case tea.WindowSizeMsg:
		m.topPane.SetSize(msg.Width, 4)
		m.leftPane.SetSize(msg.Width/2, msg.Height-4)
		m.rightPane.SetSize(msg.Width/2, msg.Height-4)

		if !m.ready {
			m.ready = true
		}
		return m, nil
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "exit"))):
			return m, tea.Quit

		case key.Matches(msg, key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Scroll up"))):
			m.list.CursorUp()
			m.leftPane.SetContent(m.list.View())

		case key.Matches(msg, key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Scroll down"))):
			m.list.CursorDown()
			m.BuildLeftPane()

		case key.Matches(msg, key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "Select"))):
			m.handleEnterKey()
			m.BuildLeftPane()
			m.BuildRightPane()
		}
	}

	return m, nil
}
