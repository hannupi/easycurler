package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	selectedBorderColor = "61"
)

var reqMethodStyleBorder = lipgloss.Border{
	Top:         lipgloss.NormalBorder().Top,
	Bottom:      lipgloss.NormalBorder().Bottom,
	Left:        lipgloss.RoundedBorder().Left,
	Right:       lipgloss.NormalBorder().Right,
	TopLeft:     lipgloss.RoundedBorder().TopLeft,
	TopRight:    lipgloss.NormalBorder().TopRight,
	BottomLeft:  lipgloss.RoundedBorder().BottomLeft,
	BottomRight: lipgloss.NormalBorder().BottomRight,
}

func (m model) View() string {
	reqMethodStyle := lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center).
		Border(reqMethodStyleBorder).
		Width(10)
	if m.FocusedComponent == FocusReqMethod {
		reqMethodStyle = reqMethodStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	urlInputStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true, true, true, false).
		Padding(0, 1).
		Width(m.Width - 16)
	if m.FocusedComponent == FocusURL {
		urlInputStyle = urlInputStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	responseBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Width(m.Width - 5).
		Height(m.Height - 10)
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
