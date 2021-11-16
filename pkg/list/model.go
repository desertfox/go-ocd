package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Kind          string
	Height        int
	Width         int
	list          list.Model
	items         []list.Item
	selected      bool
	style         lipgloss.Style
	selectedStyle lipgloss.Style
}

func NewModel(kind string, height int, width int, newItems []string, style lipgloss.Style, selectedStyle lipgloss.Style) Model {
	m := Model{}

	m.SetKind(kind)

	m.AddItems(newItems)

	list := list.NewModel(m.items, newItemDelegate(), 0, 0)
	list.SetShowTitle(false)
	m.list = list

	m.SetSize(height, width)

	m.selected = true

	m.style = style
	m.selectedStyle = selectedStyle

	return m
}

func (m *Model) SetSize(height, width int) {
	m.Width = width
	m.Height = height
	m.list.SetSize(width, height)
}

func (m *Model) SetKind(n string) {
	m.Kind = n
}

func (m *Model) ToggleSelected() {
	m.selected = !m.selected
}
