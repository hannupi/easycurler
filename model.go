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
	output           string
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
	ti.Prompt = "URL: "
	ti.CharLimit = 2048
	ti.Width = 40
	ti.SetValue("example.com")
	ti.Focus()

	vp := viewport.New(80, 20)

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
		case "tab":
			m.focusedComponent = (m.focusedComponent + 1) % NumOfFocusableComponents
			if m.focusedComponent == focusURL {
				m.urlInput.Focus()
			} else {
				m.urlInput.Blur()
			}
			return m, nil
		case "enter":
			return m, fetchURL(m.urlInput.Value())
		}
		m.urlInput, cmd = m.urlInput.Update(msg)
		return m, cmd

	case httpResMsg:
		m.output = string(msg)
		return m, nil

	case errMsg:
		m.err = msg
		m.output = "Error: " + msg.Error()
		return m, nil
	}
	return m, nil
}
