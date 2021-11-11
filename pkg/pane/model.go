package pane

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	viewport viewport.Model
	style    lipgloss.Style
	width    int
	height   int
}

func NewModel(content string) Model {
	m := Model{}

	m.SetContent(content)

	return m
}

func (m *Model) SetContent(content string) {
	m.viewport.SetContent(content)
}

func (m *Model) Width() int {
	return m.viewport.Width
}

func (m *Model) Height() int {
	return m.viewport.Height
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height

	m.style = lipgloss.NewStyle().
		Width(m.width).
		Height(m.height) /*.
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#fe8019"))*/

	m.viewport.Width = width - m.style.GetHorizontalBorderSize()
	m.viewport.Height = height - m.style.GetVerticalBorderSize()
}
