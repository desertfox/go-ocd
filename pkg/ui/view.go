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

func (m *Model) buildPane(p buildPaneMsg) (tea.Model, tea.Cmd) {
	switch p {
	case "left":
		m.buildLeftPane()
	case "right":
		m.buildRightPane()
	case "top":
		m.buildTopPane()
	}
	return m, nil
}

func (m *Model) buildTopPane() {
	msg := lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render(fmt.Sprintf("Namespace: %s", m.namespace))
	m.topPane.SetContent(msg)
}

func (m *Model) buildLeftPane() {
	m.leftPane.SetContent(m.list.View())
}

func (m *Model) buildRightPane() {
	help := m.help.View(m.keys)

	m.helpStyle = m.helpStyle.PaddingTop(m.rightPane.Height() - lipgloss.Height(help)).PaddingLeft(m.rightPane.Width() - lipgloss.Width(help))

	m.rightPane.SetContent(m.rightPaneContent + m.helpStyle.Render(help))
}
