package list

func (m *Model) CursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *Model) CursorDown() {
	if m.cursor < len(m.items)-1 {
		m.cursor++
	}
}

func (m *Model) CursorReset() {
	m.cursor = 0
}

func (m *Model) GetCursor() int {
	return m.cursor
}
