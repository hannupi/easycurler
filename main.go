package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hannupi/easycurler/internal/core"
)

func main() {
	p := tea.NewProgram(core.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
