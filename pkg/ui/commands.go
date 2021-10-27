package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/list"
)

type selectNamespaceKindMsg string

type selectNamespaceMsg bool

func (m Model) selectNamespaceKindCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return selectNamespaceKindMsg(namespace)
	}
}

func (m Model) selectNamespaceCmd() tea.Cmd {
	return func() tea.Msg {
		return selectNamespaceMsg(true)
	}
}

func (m *Model) lookupNamespaceKind(namespace string) {
	//api call to get arbitrary kinds
	kind := make(map[string][]string)

	kind["BuildConfig"] = append(kind["BuildConfig"], "bcname1")
	kind["BuildConfig"] = append(kind["BuildConfig"], "bcname2")
	kind["ImageStream"] = append(kind["ImageStream"], "isname1")
	kind["ImageStream"] = append(kind["ImageStream"], "isname2")

	m.list = list.NewModel(kind)
	fmt.Printf("%#v\n", m.list)
}

func (m *Model) lookupNamespace() {
	kind := make(map[string][]string, 1)

	kind["Project"] = append(kind["Project"], "projectname2")

	m.list = list.NewModel(kind)
}
