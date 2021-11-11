package list

import (
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	kind   string
	height int
	width  int
	list   list.Model
}

func NewModel(kind string, newItems []string) Model {
	m := Model{}

	m.SetKind(kind)

	m.AddItems(newItems)

	return m
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.list.SetSize(width, height)
}

func (m *Model) SetKind(n string) {
	m.kind = n
}
