package list

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var items []string

	items = append(items, m.headingStyle.Copy().Render(fmt.Sprintf("Kind: %s\n", m.kind)))

	for i, name := range m.items {
		if m.cursor == i {
			items = append(items, m.selectedStyle.Render(fmt.Sprintf("%s\n", name)))
		} else {
			items = append(items, m.style.Copy().Render(fmt.Sprintf("%s\n", name)))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Top, items...)
}
