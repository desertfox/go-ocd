package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-ocd/pkg/list"
	"github.com/go-ocd/pkg/pane"
)

type kind string
type namespace string

type Model struct {
	topPane pane.Model

	leftPane  pane.Model
	list      list.Model
	namespace namespace
	kind      kind

	rightPane        pane.Model
	rightPaneContent string
	help             help.Model
	helpStyle        lipgloss.Style

	keys  keyMap
	ready bool
}

func NewModel() Model {
	m := Model{}

	m.topPane = pane.NewModel()
	m.leftPane = pane.NewModel()
	m.rightPane = pane.NewModel()

	m.help = help.NewModel()
	m.helpStyle = lipgloss.NewStyle()

	m.keys = keys

	m.ready = false

	m.namespace = namespace("")
	m.kind = kind("")

	return m
}

func (m *Model) SetNamespace(n namespace) {
	m.namespace = n
}

func (m *Model) SetKind(k kind) {
	m.kind = k
}

func (m *Model) GetKind() kind {
	return m.kind
}

func (m Model) GetNamespaces() []string {
	return []string{"namespace1", "namespace2"}
}

func (m Model) GetBuildConfig() []string {
	return []string{"bc1", "bc2"}
}
