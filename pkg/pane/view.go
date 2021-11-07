package pane

func (m Model) View() string {
	return m.style.Copy().
		Width(m.viewport.Width).
		Render(m.viewport.View())
}
