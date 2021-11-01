package list

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var items []string

	items = append(items, m.headingStyle.Copy().Render(fmt.Sprintf("Kind: %s", m.kind)))

	for i, name := range m.items {
		if m.cursor == i {
			items = append(items, m.selectedStyle.Padding(1, 0, 1, 0).Render(string(name)))
		} else {
			items = append(items, m.style.Copy().Render(string(name)))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Top, items...)
}
