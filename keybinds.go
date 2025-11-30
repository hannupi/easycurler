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
		if m.FocusedComponent != FocusURL {
			return m, tea.Quit
		}
	case "down", "j":
		if m.FocusedComponent == FocusResponseBox {
			m.ResponseBox.LineDown(3)
			return m, nil
		}
	case "up", "k":
		if m.FocusedComponent == FocusResponseBox {
			m.ResponseBox.LineUp(3)
			return m, nil
		}
	case "tab":
		m.FocusedComponent = (m.FocusedComponent + 1) % NumOfFocusableComponents
		if m.FocusedComponent == FocusURL {
			m.URLInput.Focus()
		} else {
			m.URLInput.Blur()
		}
		return m, nil
	case "shift+tab":
		m.FocusedComponent = (m.FocusedComponent - 1 + NumOfFocusableComponents) % NumOfFocusableComponents
		if m.FocusedComponent == FocusURL {
			m.URLInput.Focus()
		} else {
			m.URLInput.Blur()
		}
		return m, nil
	case "enter":
		if m.FocusedComponent == FocusURL {
			return m, fetchURL(m.URLInput.Value(), m.SelectedMethod.String())
		}
		if m.FocusedComponent == FocusReqMethod {
			m.DropDownOpen = true
			return m, nil
		}
	case "?":
		// TODO help keybind menu
		return m, nil
	}

	if m.FocusedComponent == FocusURL {
		m.URLInput, cmd = m.URLInput.Update(msg)
	}
	return m, cmd
}
