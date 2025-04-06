package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type focusableComponent int
type reqMethod string

func (r reqMethod) FilterValue() string { return string(r) }
func (r reqMethod) String() string      { return string(r) }

type methodDelegate struct{}

func (d methodDelegate) Height() int  { return 1 }
func (d methodDelegate) Spacing() int { return 0 }
func (d methodDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}
func (d methodDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	method := item.(reqMethod)
	cursor := "  "
	if index == m.Index() {
		cursor = "> "
	}
	fmt.Fprintf(w, "%s%s", cursor, method)
}

const (
	FocusURL focusableComponent = iota
	FocusViewport
	FocusReqMethod
	NumOfFocusableComponents
)

type model struct {
	ReqMethods       list.Model
	SelectedMethod   reqMethod
	DropDownOpen     bool
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
	methods := []list.Item{
		reqMethod("GET"),
		reqMethod("POST"),
		reqMethod("PUT"),
		reqMethod("DELETE"),
		reqMethod("PATCH"),
	}

	const width = 20
	var height = len(methods) + 2
	l := list.New(methods, methodDelegate{}, width, height)
	l.SetShowTitle(false)
	l.SetShowPagination(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
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
		ReqMethods:       l,
		SelectedMethod:   reqMethod("GET"),
		DropDownOpen:     false,
	}
}

type httpResMsg string
type errMsg error

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
		m.Viewport.SetContent(string(msg))
		m.Viewport.GotoTop()
		return m, nil

	case errMsg:
		m.Err = msg
		return m, nil
	}
	return m, nil
}
