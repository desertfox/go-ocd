package list

import "fmt"

type Model struct {
	cursor int
	kind   string
	items  []item
}

func NewModel(kind string) Model {
	return Model{0, kind, []item{}}
}

func (m *Model) GetItemAtCursor() item {
	return m.items[m.cursor]
}

func (m *Model) AddItems(items []string) {
	var newItems []item
	for _, i := range items {
		newItems = append(newItems, newItem(i))
	}
	m.items = newItems
}

func (m Model) View() string {
	output := ""
	output += fmt.Sprintf("Kind: %s\n", m.kind)

	for i, name := range m.items {
		if m.cursor == i {
			output += fmt.Sprintf("[x] %s\n", name)
		} else {
			output += fmt.Sprintf("[ ] %s\n", name)
		}
	}

	return output
}
