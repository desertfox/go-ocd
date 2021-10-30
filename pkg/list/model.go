package list

import (
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	cursor        int
	kind          string
	style         lipgloss.Style
	selectedStyle lipgloss.Style
	headingStyle  lipgloss.Style
	items         []item
}

func NewModel(kind string) Model {
	m := Model{}

	m.cursor = 0
	m.kind = kind
	m.style = lipgloss.NewStyle()
	m.headingStyle = m.style.Copy().Foreground(lipgloss.Color("166"))
	m.selectedStyle = m.style.Copy().Foreground(lipgloss.Color("226"))

	m.items = []item{}

	return m
}
