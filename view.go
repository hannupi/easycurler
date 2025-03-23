package main

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.err != nil {
		// TODO dont error out if bad req
		return m.resContent + "\n\nPress ctrl+c to exit."
	}

	reqMethodStyle := lipgloss.NewStyle()

	urlInputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.width - 30)
	if m.focusedComponent == focusURL {
		urlInputStyle = urlInputStyle.BorderForeground(lipgloss.Color("62"))
	}

	viewportStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.width - 10).
		Height(10)
	if m.focusedComponent == focusViewport {
		viewportStyle = viewportStyle.BorderForeground(lipgloss.Color("62"))
	}

	viewportBox := viewportStyle.Render(m.viewport.View())

	urlInput := urlInputStyle.Render(m.urlInput.View())
	reqMethod := reqMethodStyle.Render(m.reqMethod.View())

	queryForms := lipgloss.JoinHorizontal(lipgloss.Bottom, reqMethod, urlInput)
	content := lipgloss.JoinVertical(lipgloss.Top, queryForms, viewportBox)

	modalStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), false, false, false, false).
		Padding(0, 1)

	modal := modalStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, modal)
}
