package ocd

import (
	"github.com/go-ocd/pkg/keys"
	"github.com/go-ocd/pkg/list"
	"github.com/go-ocd/pkg/pane"
	"github.com/go-ocd/pkg/statusbar"
	"github.com/go-ocd/pkg/styles"
)

type Model struct {
	statusbar statusbar.Model
	list      list.Model
	pane      pane.Model

	style styles.Style

	namespace string
	kind      string

	keys keys.KeyMap

	ready bool
}

func NewModel(theme string) Model {
	m := Model{}

	m.style = styles.GetTheme(theme)

	m.statusbar = statusbar.NewModel(m.style.Statusbar)

	m.list = list.NewModel("Loading", []string{})

	m.pane = pane.NewModel("", m.style.Pane)

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
