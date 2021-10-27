package list

import "fmt"

type Model struct {
	resources map[string][]string
}

func NewModel(resources map[string][]string) Model {
	fmt.Printf("list.NewModel: %#v\n", resources)
	m := Model{resources}
	return m
}

func (m *Model) SetContent(resources map[string][]string) {
	m.resources = resources
}

func (m Model) View() string {
	output := ""

	for namespace, names := range m.resources {
		output += fmt.Sprintf("%s\n", namespace)
		for _, name := range names {
			output += fmt.Sprintf("    %s\n", name)
		}
	}

	return output
}
