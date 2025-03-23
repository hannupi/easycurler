package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
