package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if !m.ready {
		return fmt.Sprintf("%s", "loading...")
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.leftPane.View(),
		m.rightPane.View(),
	)
}
