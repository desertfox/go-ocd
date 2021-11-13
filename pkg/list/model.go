package list

import (
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	Kind   string
	Height int
	Width  int
	list   list.Model
	items  []list.Item
}

func NewModel(kind string, height int, width int, newItems []string) Model {
	m := Model{}

	m.SetKind(kind)

	m.AddItems(newItems)

	list := list.NewModel(m.items, newItemDelegate(), 0, 0)
	list.SetShowTitle(false)
	m.list = list

	m.SetSize(height, width)

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
