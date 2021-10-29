package list

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	cursor    int
	kind      string
	kindStyle lipgloss.Style
	items     []item
}

func NewModel(kind string) Model {
	return Model{0, kind, lipgloss.NewStyle(), []item{}}
}

func (m Model) View() string {
	output := ""
	output += fmt.Sprintf("Kind: %s\n", m.kind)

	for i, name := range m.items {
		if m.cursor == i {
			output += fmt.Sprintf("[x] %s\n", name)
		} else {
			output += fmt.Sprintf("[ ] %s\n", name)
		}
	}

	return output
}
