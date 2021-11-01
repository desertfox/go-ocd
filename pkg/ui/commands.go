package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type setNamespaceMsg string
type getNamespacesMsg []string
type setKindMsg string
type buildPaneMsg string
type getKindInstanceDescribeMsg string

func (m Model) setNamespaceCmd(namespace string) tea.Cmd {
	return func() tea.Msg {
		return setNamespaceMsg(namespace)
	}
}

func (m Model) getNamespacesCmd() tea.Cmd {
	fakeNamespaces := m.GetNamespaces()

	return func() tea.Msg {
		return getNamespacesMsg(fakeNamespaces)
	}
}

func (m Model) setKindCmd(kind string) tea.Cmd {
	return func() tea.Msg {
		return setKindMsg(kind)
	}
}

func (m Model) getKindInstanceDescribeCmd(instance string) tea.Cmd {
	//TODO API Service
	describe := fmt.Sprintf("%s/%s/%s", m.namespace, m.kind, instance)

	return func() tea.Msg {
		return getKindInstanceDescribeMsg(describe)
	}
}

func (m Model) batchAllPanes() tea.Cmd {
	return tea.Batch(
		m.buildPaneCmd("top"),
		m.buildPaneCmd("left"),
		m.buildPaneCmd("right"),
	)
}

func (m Model) buildPaneCmd(pane string) tea.Cmd {
	return func() tea.Msg {
		return buildPaneMsg(pane)
	}
}
