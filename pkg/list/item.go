package list

type item string

func newItem(name string) item {
	return item(name)
}

func (m *Model) GetItemAtCursor() string {
	return string(m.items[m.cursor])
}

func (m *Model) AddItems(items []string) {
	var newItems []item
	for _, i := range items {
		newItems = append(newItems, newItem(i))
	}
	m.items = newItems
}
