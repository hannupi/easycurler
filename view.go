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
		Border(reqMethodStyleBorder).
		Width(m.ReqMethods.Width()).
		AlignHorizontal(lipgloss.Center)
	if m.FocusedComponent == FocusReqMethod {
		reqMethodStyle = reqMethodStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	urlInputStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true, true, true, false).
		Padding(0, 1).
		Width(m.URLInput.Width)
	if m.FocusedComponent == FocusURL {
		urlInputStyle = urlInputStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	requestSettingsBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Width(m.RequestSettingsBox.Width).
		Height(m.RequestSettingsBox.Height)
	if m.FocusedComponent == FocusRequestSettingsBox {
		requestSettingsBoxStyle = requestSettingsBoxStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	responseBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Width(m.ResponseBox.Width).
		Height(m.ResponseBox.Height)
	if m.FocusedComponent == FocusResponseBox {
		responseBoxStyle = responseBoxStyle.BorderForeground(lipgloss.Color(selectedBorderColor))
	}

	requestSettingsBox := requestSettingsBoxStyle.Render(m.RequestSettingsBox.View())

	responseBox := responseBoxStyle.Render(m.ResponseBox.View())
	if m.DropDownOpen {
		return fmt.Sprintf("Select HTTP Method:\n%s\n(Enter to select next, Esc to cancel)", m.ReqMethods.View())
	}

	urlInput := urlInputStyle.Render(m.URLInput.View())
	reqMethod := reqMethodStyle.Render(string(m.SelectedMethod))

	queryForms := lipgloss.JoinHorizontal(lipgloss.Bottom, reqMethod, urlInput)
	content := lipgloss.JoinVertical(lipgloss.Top, queryForms, requestSettingsBox, responseBox)

	modalStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), false, false, false, false).
		Padding(0, 1)

	modal := modalStyle.Render(content)

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, modal)
}
