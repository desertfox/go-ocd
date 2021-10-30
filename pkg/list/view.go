package list

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var items []string

	items = append(items, m.style.Copy().
		Foreground(lipgloss.Color("#FFF7DB")).
		Render(fmt.Sprintf("Kind: %s\n", m.kind)))

	for i, name := range m.items {
		if m.cursor == i {
			items = append(items, m.style.Copy().Foreground(lipgloss.Color("#FFFF00")).Render(fmt.Sprintf("%s\n", name)))
		} else {
			items = append(items, m.style.Copy().
				Foreground(lipgloss.Color("#808080")).
				Render(fmt.Sprintf("%s\n", name)))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Top, items...)
}
