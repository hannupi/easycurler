package main

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.err != nil {
		return m.output + "\n\nPress any key to exit."
	}

	reqMethodStyle := lipgloss.NewStyle()

	inputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.width - 30)

	outputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.width - 10)

	urlInput := inputStyle.Render(m.urlInput.View())
	reqMethod := reqMethodStyle.Render(m.reqMethod.View())

	outputContent := m.output
	outputBox := outputStyle.Render(outputContent)

	queryForms := lipgloss.JoinHorizontal(lipgloss.Bottom, reqMethod, urlInput)
	content := lipgloss.JoinVertical(lipgloss.Top, queryForms, outputBox)

	modalStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), false, false, false, false).
		Padding(0, 1)

	modal := modalStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, modal)
}
