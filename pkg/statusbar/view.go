package statusbar

import (
	"fmt"
)

func (m Model) View() string {
	namespaceString := fmt.Sprintf("Namespace: %s", m.namespace)
	kindString := fmt.Sprintf("Kind: %s", m.kind)

	output := fmt.Sprintf("%s\n%s", namespaceString, kindString)

	return m.
		style.
		Width(m.Dimension.W).
		Height(m.Dimension.H).
		Render(output)
}
