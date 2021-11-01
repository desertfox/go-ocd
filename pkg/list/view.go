package list

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var items []string

	for i, name := range m.items {
		if m.cursor == i {
			items = append(items, m.selectedStyle.Render(string(name)))
		} else {
			items = append(items, m.style.Render(string(name)))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Top, items...)
}
