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

func NewModel(style lipgloss.Style, selectedStyle lipgloss.Style) Model {
	m := Model{}

	m.AddItems([]string{})

	list := list.NewModel(m.items, newItemDelegate(), 0, 0)
	list.SetShowTitle(false)
	m.list = list

	m.selected = true

	m.style = style
	m.selectedStyle = selectedStyle

	return m
}

func (m *Model) SetNewList(kind string, newItems []string) {
	m.SetKind(kind)

	m.AddItems(newItems)

	list := list.NewModel(m.items, newItemDelegate(), 0, 0)
	list.SetShowTitle(false)
	m.list = list

	m.SetSize(m.Dimension)

}

func (m *Model) SetSize(d helpers.Dimension) {
	m.Dimension = d

	m.list.SetSize(m.Dimension.W, m.Dimension.H)
}

func (m *Model) SetKind(n string) {
	m.Kind = n
}

func (m *Model) ToggleSelected() {
	m.selected = !m.selected
}

func (m *Model) IsFiltering() bool {
	return m.list.FilterState() == list.Filtering
}
