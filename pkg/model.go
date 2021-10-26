package ui

import (
	//"github.com/go-ocd/pkg/list"
	"github.com/go-ocd/pkg/pane"
)

type Model struct {
	leftPane  pane.Model
	rightPane pane.Model
	ready     bool
	//resources list.Model
}

func NewModel() Model {
	leftPane := pane.NewModel()
	rightPane := pane.NewModel()

	//resources := []list.Resource{list.NewResource("a", "b", "c")}

	return Model{leftPane, rightPane, false}
}
