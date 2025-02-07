package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	input  textinput.Model
	output string
	err    error
	width  int
	height int
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter URL"
	ti.Prompt = "URL: "
	ti.CharLimit = 256
	ti.Width = 40
	ti.SetValue("https://example.com")
	ti.Focus()

	return model{
		input:  ti,
		output: "Curl output will appear here...",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func fetchURL(url string) tea.Cmd {
	return func() tea.Msg {
		out, err := exec.Command("curl", "-s", url).CombinedOutput()
		if err != nil {
			return errMsg(err)
		}
		return curlOutputMsg(strings.TrimSpace(string(out)))
	}
}

type curlOutputMsg string
type errMsg error

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, fetchURL(m.input.Value())
		}
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd

	case curlOutputMsg:
		m.output = string(msg)
		return m, nil

	case errMsg:
		m.err = msg
		m.output = "Error: " + msg.Error()
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return m.output + "\n\nPress any key to exit."
	}

	inputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Width(m.width - 10)

	outputStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Width(m.width - 10)

	inputBox := inputStyle.Render(m.input.View())

	outputContent := fmt.Sprintf("Curl Output:\n\n%s\n\nPress Enter to fetch, q or Ctrl+C to quit.", m.output)
	outputBox := outputStyle.Render(outputContent)

	content := lipgloss.JoinVertical(lipgloss.Top, inputBox, outputBox)

	modalStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)

	modal := modalStyle.Render(content)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, modal)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

