package statusbar

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	namespaceString := fmt.Sprintf("Namespace: %s", m.namespace)
	kindString := fmt.Sprintf("Kind: %s", m.kind)

	return m.
		style.
		Padding(1, 1, 1, 1).
		Bold(true).
		Foreground(lipgloss.Color("#fe8019")).
		Width(m.width).
		Height(m.height).
		Render(fmt.Sprintf("%s\n%s", namespaceString, kindString))
}
