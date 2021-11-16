package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var themes map[string]Style

type Style struct {
	Statusbar lipgloss.Style
	List      lipgloss.Style
	Pane      lipgloss.Style
	Selected  lipgloss.Style
}

func init() {
	themes = make(map[string]Style, 1)

	themes["halloween"] = newHalloweenTheme()
}

func GetTheme(theme string) Style {
	return themes[theme]
}

func newHalloweenTheme() Style {
	return Style{
		Statusbar: lipgloss.NewStyle().Padding(1, 1, 1, 1).Foreground(lipgloss.Color("#fe8019")),
		List:      lipgloss.NewStyle(),
		Pane:      lipgloss.NewStyle(),
		Selected: lipgloss.NewStyle().Padding(2, 2, 2, 2).Border(lipgloss.NormalBorder()).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#fe8019")),
	}
}
