package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-ocd/pkg/helpers"
)

type Model struct {
	Kind          string
	Dimension     helpers.Dimension
	list          list.Model
	items         []list.Item
	selected      bool
	style         lipgloss.Style
	selectedStyle lipgloss.Style
}

func NewModel(kind string, d helpers.Dimension, newItems []string, style lipgloss.Style, selectedStyle lipgloss.Style) Model {
	m := Model{}

	m.SetKind(kind)

	m.AddItems(newItems)

	list := list.NewModel(m.items, newItemDelegate(), 0, 0)
	list.SetShowTitle(false)
	m.list = list

	m.Dimension = d

	m.SetSize(d.W, d.H)

	m.selected = true

	m.style = style
	m.selectedStyle = selectedStyle

	return m
}

func (m *Model) SetSize(w, h int) {
	m.Dimension.Set(w, h)

	m.list.SetSize(w, h)
}

func (m *Model) SetKind(n string) {
	m.Kind = n
}

func (m *Model) ToggleSelected() {
	m.selected = !m.selected
}
