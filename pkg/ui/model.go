package ui

import (
	"github.com/go-ocd/pkg/list"
	"github.com/go-ocd/pkg/pane"
)

type kind string
type namespace string

type Model struct {
	topPane   pane.Model
	leftPane  pane.Model
	rightPane pane.Model
	list      list.Model

	namespace namespace
	kind      kind

	ready bool
}

func NewModel() Model {
	m := Model{}

	m.topPane = pane.NewModel()
	m.leftPane = pane.NewModel()
	m.rightPane = pane.NewModel()

	m.ready = false

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
