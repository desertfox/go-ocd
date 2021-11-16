package pane

func (m Model) View() string {

	if m.selected {
		return m.selectedStyle.Width(m.viewport.Width).Render(m.viewport.View())
	}

	return m.style.Copy().
		Width(m.viewport.Width).
		Render(m.viewport.View())
}
