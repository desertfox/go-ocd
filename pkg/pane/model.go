package pane

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-ocd/pkg/helpers"
)

type Model struct {
	viewport      viewport.Model
	Dimension     helpers.Dimension
	selected      bool
	style         lipgloss.Style
	selectedStyle lipgloss.Style
}

func NewModel(content string, style lipgloss.Style, selectedStyle lipgloss.Style) Model {
	m := Model{}

	m.SetContent(content)

	m.style = style
	m.selectedStyle = selectedStyle

	m.selected = false

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

func (m *Model) SetSize(d helpers.Dimension) {
	m.Dimension = d

	m.style = m.style.Copy().
		Width(m.Dimension.W).
		Height(m.Dimension.H)

	m.viewport.Width = m.Dimension.W - m.style.GetHorizontalBorderSize()
	m.viewport.Height = m.Dimension.H - m.style.GetVerticalBorderSize()
}

func (m *Model) ToggleSelected() {
	m.selected = !m.selected
}
