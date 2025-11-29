package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	selectedBorderColor = "61"
)

func (m model) View() string {
	if m.Err != nil {
		// TODO recover from bad req
		m.ResponseBox.SetContent(m.Err.Error())
	}

	reqMethodStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1)
	if m.FocusedComponent == FocusReqMethod {
		reqMethodStyle = reqMethodStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	urlInputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.Width - 30)
	if m.FocusedComponent == FocusURL {
		urlInputStyle = urlInputStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	responseBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Width(m.Width - 10).
		Height(m.Height - 30)
	if m.FocusedComponent == FocusResponseBox {
		responseBoxStyle = responseBoxStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	responseBox := responseBoxStyle.Render(m.ResponseBox.View())
	if m.DropDownOpen {
		return fmt.Sprintf("Select HTTP Method:\n%s\n(Enter to select next, Esc to cancel)", m.ReqMethods.View())
	}

	urlInput := urlInputStyle.Render(m.UrlInput.View())
	reqMethod := reqMethodStyle.Render(string(m.SelectedMethod))

	queryForms := lipgloss.JoinHorizontal(lipgloss.Bottom, reqMethod, urlInput)
	content := lipgloss.JoinVertical(lipgloss.Top, queryForms, responseBox)

	modalStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), false, false, false, false).
		Padding(0, 1)

	modal := modalStyle.Render(content)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, modal)
}
