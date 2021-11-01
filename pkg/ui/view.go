package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if !m.ready {
		return "loading..."
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Top,
			m.panes["selected"].View(),
			m.panes["list"].View(),
			m.panes["help"].View(),
		),
		m.panes["preview"].View(),
	)
}

func (m *Model) buildPane(p buildPaneMsg) (tea.Model, tea.Cmd) {
	switch p {
	case "list":
		m.panes["list"].SetContent(m.list.View())
	case "preview":
		m.panes["preview"].SetContent(m.preview)
	case "selected":
		msg := lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render(fmt.Sprintf("Namespace: %s", m.namespace)),
			lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Render(fmt.Sprintf("Kind     : %s", m.kind)),
		)
		m.panes["selected"].SetContent(msg)
	case "help":
		m.panes["help"].SetContent(m.help.View(m.keys))
	}

	return m, nil
}
