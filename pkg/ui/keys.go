package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Help   key.Binding
	Quit   key.Binding
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Delete key.Binding
	Create key.Binding
	Edit   key.Binding
	Back   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Back, k.Enter},
		{k.Delete, k.Create, k.Edit, k.Help, k.Quit},
	}
}

var keys = keyMap{
	Help: key.NewBinding(
		key.WithKeys("h", "?"),
		key.WithHelp("h/?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),

	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),

	Back: key.NewBinding(
		key.WithKeys("left", "b"),
		key.WithHelp("←/b", "go back"),
	),

	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("<enter>", "select"),
	),

	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "create"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
}
