package list

func (m Model) View() string {
	if m.selected {
		return m.selectedStyle.Render(m.list.View())
	}
	return m.style.Render(m.list.View())
}
