package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type focusableComponent int

const (
	focusURL focusableComponent = iota
	focusViewport
	NumOfFocusableComponents
)

type model struct {
	reqMethod        list.Model
	urlInput         textinput.Model
	resContent       string
	err              error
	width            int
	height           int
	viewport         viewport.Model
	focusedComponent focusableComponent
}

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter URL"
	ti.CharLimit = 2048
	ti.Width = 40
	ti.SetValue("example.com")
	ti.Focus()

	vp := viewport.New(50, 10)
	vp.SetContent("asdfasdf")

	return model{
		urlInput:         ti,
		viewport:         vp,
		focusedComponent: 0,
	}
}

type httpResMsg string
type errMsg error

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
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
			return nil, nil
		}
		m.urlInput, cmd = m.urlInput.Update(msg)
		return m, cmd

	case httpResMsg:
		m.viewport.SetContent(string(msg))
		m.viewport.GotoTop()
		return m, nil

	case errMsg:
		m.err = msg
		m.resContent = "Error: " + msg.Error()
		return m, nil
	}
	return m, nil
}
