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
	m := Model{}
	m.SetContent("")
	return m
}

func (m *Model) SetContent(content string) {
	m.viewport.SetContent(content)
}

func (m *Model) SetSize(width, height int) {
	border := lipgloss.NormalBorder()

	m.style = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
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
