package statusbar

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/go-ocd/pkg/helpers"
)

type Model struct {
	namespace string
	kind      string
	Dimension helpers.Dimension
	style     lipgloss.Style
}

func NewModel(style lipgloss.Style) Model {
	return Model{"", "", helpers.NewDimensions(), style}
}

func (m *Model) SetNamespace(n string) {
	m.namespace = n
}

func (m *Model) SetKind(n string) {
	m.kind = n
}

func (m *Model) SetSize(w, h int) {
	m.Dimension.Set(w, h)
}
