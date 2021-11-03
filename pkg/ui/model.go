package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-ocd/pkg/list"
	"github.com/go-ocd/pkg/pane"
)

type kind string
type namespace string

type Model struct {
	panes map[string]*pane.Model

	list    list.Model
	preview string

	namespace namespace
	kind      kind

	help      help.Model
	helpStyle lipgloss.Style

	keys  keyMap
	ready bool
}

func NewModel() Model {
	m := Model{}

	m.panes = map[string]*pane.Model{
		"selected": pane.NewModel(""),
		"list":     pane.NewModel(""),
		"help":     pane.NewModel(""),
		"preview":  pane.NewModel(""),
	}

	m.help = help.NewModel()
	m.helpStyle = lipgloss.NewStyle()

	m.keys = keys

	m.ready = false

	return m
}

func (m Model) GetNamespaces() []string {
	var namespaces []string
	for i := 1; i <= 100; i++ {
		namespaces = append(namespaces, fmt.Sprintf("namespace%d", i))
	}
	return namespaces
}

func (m Model) GetBuildConfig() []string {
	return []string{"bc1", "bc2"}
}
