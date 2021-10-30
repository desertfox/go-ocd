package pane

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	viewport viewport.Model
	style    lipgloss.Style
}

func NewModel(s lipgloss.Style) Model {
	m := Model{}

	m.SetContent("")

	m.style = s

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

	m.viewport.Width = width - m.style.GetHorizontalBorderSize()
	m.viewport.Height = height - m.style.GetVerticalBorderSize()
}

func (m Model) View() string {
	return m.style.Copy().
		Width(m.viewport.Width).
		Render(m.viewport.View())
}
