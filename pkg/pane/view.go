package pane

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	return m.style.Copy().
		BorderForeground(lipgloss.Color("#b8bb26")).
		Border(lipgloss.NormalBorder()).
		Width(m.viewport.Width).
		Render(m.viewport.View())
}
