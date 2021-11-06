package statusbar

import (
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	namespace string
	kind      string
	width     int
	height    int
	style     lipgloss.Style
}

func NewModel() Model {
	return Model{"", "", 0, 0, lipgloss.NewStyle()}
}

func (m *Model) SetNamespace(n string) {
	m.namespace = n
}

func (m *Model) SetKind(n string) {
	m.kind = n
}

func (m *Model) SetSize(w int, h int) {
	m.width = w
	m.height = h
}
