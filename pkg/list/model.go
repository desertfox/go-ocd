package list

import "fmt"

var defaultKinds = []string{"BuildConfig", "ImageStream", "DeploymentConfig"}

type Model struct {
	cursor    int
	kind      string
	namespace string
	resources []string
}

func NewModel() Model {
	return Model{0, "", "", defaultKinds}
}

func (m *Model) SetNamespace(namespace string) {
	m.namespace = namespace
}

func (m *Model) SetKind(kind string) {
	m.kind = kind
}

func (m *Model) SetResources(resources []string) {
	m.resources = resources
}

func (m *Model) SetDefaultKind() {
	m.SetKind("")
	m.resources = defaultKinds
}

func (m Model) View() string {
	output := ""
	output += fmt.Sprintf("Namespace:%s\n", m.namespace)

	var tempKind string
	if m.kind == "" {
		tempKind = m.resources[m.GetCursor()]
	} else {
		tempKind = m.kind
	}
	output += fmt.Sprintf("Kind:%s\n", tempKind)

	for i, resource := range m.resources {
		if m.cursor == i {
			output += fmt.Sprintf("[x] %s\n", resource)
		} else {
			output += fmt.Sprintf("[ ] %s\n", resource)
		}
	}

	return output
}
