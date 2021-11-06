package list

import "github.com/charmbracelet/bubbles/list"

type item string

func (i item) FilterValue() string { return string(i) }
func (i item) Title() string       { return string(i) }
func (i item) Description() string { return string(i) }

func newItem(name string) item {
	return item(name)
}

func (m *Model) GetItemAtCursor() string {
	item := m.list.SelectedItem().(item)

	return item.Title()
}

func (m *Model) AddItems(newItems []string) {
	var (
		delegateKeys = newDelegateKeyMap()
	)

	items := make([]list.Item, len(newItems))
	for i := 0; i < len(newItems); i++ {
		items[i] = newItem(newItems[i])
	}

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	m.list = list.NewModel(items, delegate, 0, 0)
}
