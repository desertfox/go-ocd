package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/go-ocd/pkg/list"
)

func (m Model) View() string {
	if !m.ready {
		return fmt.Sprintf("%s", "loading...")
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.topPane.View(),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.leftPane.View(),
			m.rightPane.View(),
		),
	)
}

func (m *Model) BuildNamespacesList() {
	m.list = list.NewModel("namespace")
	m.list.AddItems(m.GetNamespaces())
}

func (m *Model) BuildKindsList() {
	var defaultKinds = []string{"BuildConfig", "ImageStream", "DeploymentConfig"}

	m.list = list.NewModel("kinds")
	m.list.AddItems(defaultKinds)
}

func (m *Model) handleEnterKey() {
	m.SetKindFromList()
	m.BuildListFromKind()
}

func (m *Model) SetKindFromList() {
	switch m.kind {
	case "namespace":
		m.SetKind("kinds")
		m.SetNamespace(string(m.list.GetItemAtCursor()))
	case "kinds":
		m.SetKind(string(m.list.GetItemAtCursor()))
	}
}

func (m *Model) BuildListFromKind() {
	switch m.kind {
	case "namespace":
		m.BuildNamespacesList()
	case "kinds":
		m.BuildKindsList()
	default:
		m.list = list.NewModel(string(m.GetKind()))
		m.list.AddItems(m.GetBuildConfig())
	}
}

func (m *Model) BuildLeftPane() {
	m.leftPane.SetContent(m.list.View())
}
func (m *Model) BuildRightPane() {
	if m.kind == "namespace" || m.kind == "kinds" {
		m.rightPane.SetContent(fmt.Sprintf("Please select %s", m.kind))
	} else {
		m.rightPane.SetContent("describe todo")
	}
}
