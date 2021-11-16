package ocd

import (
	"github.com/go-ocd/pkg/api"
	"github.com/go-ocd/pkg/helpers"
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

	keys *keys.MasterKeyMap

	ready        bool
	selectedView string

	api api.Client
}

func NewModel(theme string, kubeconfig string, fake bool) Model {
	m := Model{}

	m.style = styles.GetTheme(theme)

	m.statusbar = statusbar.NewModel(m.style.Statusbar)

	m.list = list.NewModel("Loading", helpers.NewDimensions(), []string{}, m.style.List, m.style.Selected)

	m.pane = pane.NewModel("", m.style.Pane, m.style.Selected)

	m.keys = keys.Keys

	m.ready = false

	m.selectedView = "list"

	m.api = api.NewClient(kubeconfig, fake)

	return m
}

func (m *Model) SetNamespace(n string) {
	m.namespace = n
}

func (m *Model) SetKind(n string) {
	m.kind = n
}
