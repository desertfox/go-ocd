package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/ui"
)

func main() {
	var opts []tea.ProgramOption
	opts = append(opts, tea.WithAltScreen())

	m := ui.NewModel()
	p := tea.NewProgram(m, opts...)

	if err := p.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
