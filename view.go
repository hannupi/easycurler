package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.Err != nil {
		// TODO recover from bad req
		m.Viewport.SetContent(m.Err.Error())
	}

	reqMethodStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1)
	if m.FocusedComponent == FocusReqMethod {
		reqMethodStyle = reqMethodStyle.BorderForeground(lipgloss.Color("62"))
	}

	urlInputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.Width - 30)
	if m.FocusedComponent == FocusURL {
		urlInputStyle = urlInputStyle.BorderForeground(lipgloss.Color("62"))
	}

	viewportStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.Width - 10).
		Height(10)
	if m.FocusedComponent == FocusViewport {
		viewportStyle = viewportStyle.BorderForeground(lipgloss.Color("62"))
	}

	viewportBox := viewportStyle.Render(m.Viewport.View())
	if m.DropDownOpen {
		return fmt.Sprintf("Select HTTP Method:\n%s\n(Enter to select next, Esc to cancel)", m.ReqMethods.View())
	}

	urlInput := urlInputStyle.Render(m.UrlInput.View())
	reqMethod := reqMethodStyle.Render(string(m.SelectedMethod))

	queryForms := lipgloss.JoinHorizontal(lipgloss.Bottom, reqMethod, urlInput)
	content := lipgloss.JoinVertical(lipgloss.Top, queryForms, viewportBox)

	modalStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), false, false, false, false).
		Padding(0, 1)

	modal := modalStyle.Render(content)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, modal)
}
