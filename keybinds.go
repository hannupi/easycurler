package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func handleKeyInput(msg tea.KeyMsg, m model) (model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "q":
		if m.focusedComponent != focusURL {
			return m, tea.Quit
		}
	case "j":
		if m.focusedComponent == focusViewport {
			m.viewport.LineDown(3)
			return m, nil
		}
	case "k":
		if m.focusedComponent == focusViewport {
			m.viewport.LineUp(3)
			return m, nil
		}
	case "tab":
		m.focusedComponent = (m.focusedComponent + 1) % NumOfFocusableComponents
		if m.focusedComponent == focusURL {
			m.urlInput.Focus()
		} else {
			m.urlInput.Blur()
		}
		return m, nil
	case "shift+tab":
		m.focusedComponent = (m.focusedComponent - 1) % NumOfFocusableComponents
		if m.focusedComponent == focusURL {
			m.urlInput.Focus()
		} else {
			m.urlInput.Blur()
		}
		return m, nil
	case "enter":
		if m.focusedComponent == focusURL {
			return m, fetchURL(m.urlInput.Value())
		}
	case "?":
		// TODO help keybind menu
		return m, nil
	}

	if m.focusedComponent == focusURL {
		m.urlInput, cmd = m.urlInput.Update(msg)
	}
	return m, cmd
}
