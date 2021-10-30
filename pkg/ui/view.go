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
	msg := fmt.Sprintf("Namespace: %s\n", m.namespace)
	m.topPane.SetContent(msg)
}

func (m *Model) buildLeftPane() {
	m.leftPane.SetContent(m.list.View())
}

func (m *Model) buildRightPane() {
	preview := fmt.Sprintf("Please select %s\n", m.kind)

	help := fmt.Sprintf("%s\n", m.help.View(m.keys))

	m.helpStyle = m.helpStyle.PaddingTop(m.rightPane.Height() - lipgloss.Height(help)).PaddingLeft(m.rightPane.Width() - lipgloss.Width(help))

	m.rightPane.SetContent(preview + m.helpStyle.Render(help))
}
