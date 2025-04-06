package main

import (
	"fmt"

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
		if m.FocusedComponent == FocusViewport {
			m.Viewport.LineDown(3)
			return m, nil
		}
	case "up", "k":
		if m.FocusedComponent == FocusViewport {
			m.Viewport.LineUp(3)
			return m, nil
		}
	case "tab":
		m.FocusedComponent = (m.FocusedComponent + 1) % NumOfFocusableComponents
		if m.FocusedComponent == FocusURL {
			m.UrlInput.Focus()
		} else {
			m.UrlInput.Blur()
		}
		return m, nil
	case "shift+tab":
		m.FocusedComponent = (m.FocusedComponent - 1) % NumOfFocusableComponents
		if m.FocusedComponent == FocusURL {
			m.UrlInput.Focus()
		} else {
			m.UrlInput.Blur()
		}
		return m, nil
	case "enter":
		if m.FocusedComponent == FocusURL {
			return m, fetchURL(m.UrlInput.Value())
		}
		if m.FocusedComponent == FocusReqMethod {
			m.DropDownOpen = true
			selectedMethod := m.ReqMethods.SelectedItem()
			fmt.Println(selectedMethod)
			return m, nil
		}
	case "?":
		// TODO help keybind menu
		return m, nil
	}

	if m.FocusedComponent == FocusURL {
		m.UrlInput, cmd = m.UrlInput.Update(msg)
	}
	return m, cmd
}
