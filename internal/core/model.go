package core

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

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

func InitialModel() model {
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
	ti.Prompt = ""
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
