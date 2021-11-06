package list

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/api"
	"github.com/go-ocd/pkg/commands"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg)

	case commands.SetKindMsg:
		return m.handleSetKindMsg(msg)
	}

	return m, m.updateList(msg)
}

func (m *Model) updateList(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return cmd
}

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) {
	m.SetSize(msg.Width/2, msg.Height-20)
}

func (m *Model) handleSetKindMsg(msg commands.SetKindMsg) (Model, tea.Cmd) {
	switch msg {
	case "namespace":
		m.SetKind("namespace")
		m.AddItems(api.GetNamespaces())
	case "kind":
		m.SetKind("kind")
		m.AddItems([]string{"BuildConfig", "ImageStream", "DeploymentConfig"})
	default:
		m.SetKind(string(m.kind))
		m.AddItems(api.GetBuildConfig())
	}

	return *m, nil
}
