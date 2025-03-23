package main

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.err != nil {
		return m.output + "\n\nPress any key to exit."
	}

	reqMethodStyle := lipgloss.NewStyle()

	urlInputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.width - 30)
	if m.focusedComponent == 0 {
		urlInputStyle = urlInputStyle.BorderForeground(lipgloss.Color("62"))
	}

	viewportStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.width - 10)
	if m.focusedComponent == 1 {
		viewportStyle = viewportStyle.BorderForeground(lipgloss.Color("62"))
	}

	urlInput := urlInputStyle.Render(m.urlInput.View())
	reqMethod := reqMethodStyle.Render(m.reqMethod.View())

	outputContent := m.output
	outputBox := viewportStyle.Render(outputContent)

	queryForms := lipgloss.JoinHorizontal(lipgloss.Bottom, reqMethod, urlInput)
	content := lipgloss.JoinVertical(lipgloss.Top, queryForms, outputBox)

	modalStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), false, false, false, false).
		Padding(0, 1)

	modal := modalStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, modal)
}
