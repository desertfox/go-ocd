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

	m.kind = kind

	items := make([]list.Item, len(newItems))
	for i := 0; i < len(newItems); i++ {
		items[i] = newItem(newItems[i])
	}

	var delegateKeys = newDelegateKeyMap()
	delegate := newItemDelegate(delegateKeys)
	m.list = list.NewModel(items, delegate, 0, 0)

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
