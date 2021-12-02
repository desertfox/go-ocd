package ocd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/msgtypes"
)

func (m *Model) SetNamespaceCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return msgtypes.SetNamespaceMsg(namespace)
	}
}

func (m *Model) GetNamespacesCmd() tea.Cmd {
	namespaces := m.api.GetProjects()

	return func() tea.Msg {
		return msgtypes.GetNamespacesMsg(namespaces)
	}
}

func (m *Model) SetKindCmd(kind string) tea.Cmd {
	var items []string

	switch strings.ToLower(kind) {
	case "namespace":
		items = m.api.GetProjects()
	case "kind":
		items = m.api.GetAvailKinds()
	case "buildconfig":
		items = m.api.GetBuildConfigs(m.namespace)
	case "secrets":
		items = m.api.GetSecrets(m.namespace)
	}

	return func() tea.Msg {
		return msgtypes.SetKindMsg{Kind: kind, Items: items}
	}
}

func (m *Model) GetResourceToYAMLCmd(instance string) tea.Cmd {
	yaml := m.api.GetInstance(m.namespace, m.kind, instance)

	return func() tea.Msg {
		return msgtypes.KindInstanceYaml(yaml)
	}
}

func (m *Model) DumpToYamlCmd(instance string) tea.Cmd {
	yaml := m.api.GetInstance(m.namespace, m.kind, instance)

	filename := fmt.Sprintf("%s_%s_%s.yaml", m.namespace, m.kind, instance)

	f, err := os.Create(filename)
	if err != nil {
		return m.SignalErrorCmd(err)
	}
	defer f.Close()

	_, err = f.WriteString(yaml)
	if err != nil {
		return m.SignalErrorCmd(err)
	}

	return func() tea.Msg {
		return msgtypes.DumpToYamlMsg(true)
	}
}

func (m *Model) DeleteResourceCmd(instance string) tea.Cmd {

	kind := m.kind

	switch strings.ToLower(kind) {
	case "namespace":

	case "kind":

	case "buildconfig":
		m.api.DelBuildConfig(m.namespace, instance)
	case "secrets":

	}

	return func() tea.Msg {
		return nil
	}
}

func (m *Model) EditResourceCmd(instance string) tea.Cmd {
	yaml := m.api.GetInstance(m.namespace, m.kind, instance)

	//write to temp file
	file, err := ioutil.TempFile("tmp", "")
	if err != nil {
		return m.SignalErrorCmd(err)
	}

	file.WriteString(yaml)

	//get editor from env?

	return func() tea.Msg {
		return nil
	}

}

func (m *Model) SignalErrorCmd(err error) tea.Cmd {
	return func() tea.Msg {
		return msgtypes.Err(err)
	}
}
