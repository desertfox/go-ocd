package pane

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	viewport viewport.Model
	style    lipgloss.Style
}

func NewModel() Model {
	return Model{}
}

func (m *Model) SetContent(content string) {
	m.viewport.SetContent(content)
}

func (m *Model) SetSize(width, height int) {
	border := lipgloss.NormalBorder()
	padding := 1

	// Set the style so that the frame size is able to be determined from other components.
	m.style = lipgloss.NewStyle().
		PaddingLeft(padding).
		PaddingRight(padding).
		Border(border)

	m.viewport.Width = width - m.style.GetHorizontalBorderSize()
	m.viewport.Height = height - m.style.GetVerticalBorderSize()
}

func (m Model) View() string {
	border := lipgloss.NormalBorder()

	return m.style.Copy().
		PaddingLeft(1).
		PaddingRight(1).
		Border(border).
		Width(m.viewport.Width).
		Render(m.viewport.View())
}
