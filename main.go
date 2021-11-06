package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	ocd "github.com/go-ocd/pkg"
)

func main() {
	var opts []tea.ProgramOption
	opts = append(opts, tea.WithAltScreen())

	m := ocd.NewModel()
	p := tea.NewProgram(m, opts...)
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := p.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
