package list

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-ocd/pkg/keys"
)

func newItemDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, true).
		BorderForeground(lipgloss.Color("#fe8019")).
		Foreground(lipgloss.Color("#fe8019")).
		Padding(0, 1, 0, 1)

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		return nil
	}

	d.ShortHelpFunc = func() []key.Binding {
		return keys.Keys.ShortHelp()
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return keys.Keys.FullHelp()
	}

	d.ShowDescription = false

	return d
}
