package ocd

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	if !m.ready {
		return "loading..."
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Left,
			m.statusbar.View(),
			m.list.View(),
		),
		m.pane.View(),
	)
}
