package core

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.DropDownOpen {
		var cmd tea.Cmd
		m.ReqMethods, cmd = m.ReqMethods.Update(msg)

		if key, ok := msg.(tea.KeyMsg); ok && key.Type == tea.KeyEnter {
			m.SelectedMethod = m.ReqMethods.SelectedItem().(reqMethod)
			m.DropDownOpen = false
			return m, nil
		}
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		return handleKeyInput(msg, m)

	case httpResMsg:
		if m.Viewport.Height != m.Height-30 {
			// check if there is a cleaner way to resize the text box
			// WindowSizeMsg is sent to func Update in init?
			m.Viewport = viewport.New(0, m.Height-30)
		}
		m.Viewport.SetContent(string(msg))
		m.Viewport.GotoTop()
		return m, nil

	case errMsg:
		m.Err = msg
		return m, nil
	}
	return m, nil
}
