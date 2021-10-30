package list

import (
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	cursor int
	kind   string
	style  lipgloss.Style
	items  []item
}

func NewModel(kind string) Model {
	return Model{0, kind, lipgloss.NewStyle(), []item{}}
}
