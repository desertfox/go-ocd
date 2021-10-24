package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View returns a string representation of the entire application UI.
func (m Model) View() string {
	// If the viewport on the panes is not ready display the spinner.
	if !m.ready {
		return fmt.Sprintf("%s", "loading...")
	}

	// Return the UI with the two panes side by side and
	// the status bar at the bottom of the screen.
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.leftPane.View(),
		m.rightPane.View(),
	)
}
