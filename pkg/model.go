package ui

import (
	"github.com/go-ocd/pkg/list"
	"github.com/go-ocd/pkg/pane"
)

type Model struct {
	leftPane  pane.Model
	rightPane pane.Model
	list      list.Model
	ready     bool
}

func NewModel() Model {
	m := Model{}

	m.leftPane = pane.NewModel()
	m.rightPane = pane.NewModel()

	m.ready = false

	return m
}
