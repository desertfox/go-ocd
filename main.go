package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	ocd "github.com/go-ocd/pkg"
)

var fake bool

func init() {
	flag.BoolVar(&fake, "fake", true, "Use fake api for ui testing")
}

func main() {
	var opts []tea.ProgramOption
	opts = append(opts, tea.WithAltScreen())

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	m := ocd.NewModel("halloween", kubeconfig, fake)
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
