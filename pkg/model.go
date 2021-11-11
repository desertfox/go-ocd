package ocd

import (
	"github.com/go-ocd/pkg/keys"
	"github.com/go-ocd/pkg/list"
	"github.com/go-ocd/pkg/pane"
	"github.com/go-ocd/pkg/statusbar"
)

type Model struct {
	statusbar statusbar.Model
	list      list.Model
	pane      pane.Model

	namespace string
	kind      string

	keys keys.KeyMap

	ready bool
}

func NewModel() Model {
	m := Model{}

	m.statusbar = statusbar.NewModel()
	m.list = list.NewModel("Loading", []string{})
	m.pane = pane.NewModel("")

	m.keys = keys.Keys

	m.ready = false

	return m
}

func (m *Model) SetNamespace(n string) {
	m.namespace = n
}

func (m *Model) SetKind(n string) {
	m.kind = n
}
