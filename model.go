package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type focusableComponent int

const (
	FocusURL focusableComponent = iota
	FocusViewport
	FocusReqMethod
	NumOfFocusableComponents
)

type model struct {
	ReqMethods       list.Model
	UrlInput         textinput.Model
	Err              error
	Width            int
	Height           int
	Viewport         viewport.Model
	FocusedComponent focusableComponent
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

	return model{
		UrlInput:         ti,
		Viewport:         vp,
		FocusedComponent: 0,
	}
}

type httpResMsg string
type errMsg error

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		return handleKeyInput(msg, m)

	case httpResMsg:
		m.Viewport.SetContent(string(msg))
		m.Viewport.GotoTop()
		return m, nil

	case errMsg:
		m.Err = msg
		return m, nil
	}
	return m, nil
}
