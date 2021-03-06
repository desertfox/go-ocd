package list

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-ocd/pkg/helpers"
)

func (m Model) Update(msg tea.Msg, disableKeys bool) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg)
	}

	if !m.selected {
		return m, nil
	}

	return m, m.updateList(msg)
}

func (m *Model) updateList(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return cmd
}

func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) {
	m.SetSize(helpers.NewDimensions(msg.Width/2, msg.Height-20))
}
