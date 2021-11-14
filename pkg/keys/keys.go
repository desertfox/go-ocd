package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type MasterKeyMap struct {
	//Base
	Enter key.Binding
	Back  key.Binding
	//Instance
	Edit       key.Binding
	Delete     key.Binding
	DumpToYaml key.Binding
}

func (k MasterKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Back, k.Edit, k.Delete}
}

func (k MasterKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Enter, k.Back, k.Edit, k.Delete}, {k.DumpToYaml}}
}

func (k *MasterKeyMap) ShowInstanceKeys(show bool) {
	k.Delete.SetEnabled(show)
	k.Edit.SetEnabled(show)
	k.Back.SetEnabled(show)
}

func (k *MasterKeyMap) ShowSelectKeys(show bool) {
	k.Enter.SetEnabled(show)
}

var Keys = &MasterKeyMap{
	Back: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "go back"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("<enter>", "select"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	DumpToYaml: key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "dump to yaml file"),
	),
}
