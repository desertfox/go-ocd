package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/pane"
)

type Model struct {
	leftPane  pane.Model
	rightPane pane.Model
	ready     bool
}

func NewModel() Model {
	leftPane := pane.NewModel()
	rightPane := pane.NewModel()

	return Model{leftPane, rightPane, false}
}

func (m Model) Init() tea.Cmd {
	return nil
}
