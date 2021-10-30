package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if !m.ready {
		return fmt.Sprintf("%s", "loading...")
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.topPane.View(),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.leftPane.View(),
			m.rightPane.View(),
		),
	)
}

func (m *Model) buildTopPane() {
	m.topPane.SetContent(fmt.Sprintf("Namespace: %s\n", m.namespace))
}

func (m *Model) buildLeftPane() {
	m.leftPane.SetContent(m.list.View())
}

func (m *Model) buildRightPane() {
	preview := fmt.Sprintf("Please select %s\n", m.kind)
	help := lipgloss.NewStyle().Render(fmt.Sprintf("%s\n", m.help.View(m.keys)))
	m.rightPane.SetContent(preview + help)
}
