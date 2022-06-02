package main

import (
	"charming-todo/ui"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, _ := tea.LogToFile("debug.log", "")

	p := tea.NewProgram(ui.New(), tea.WithAltScreen())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}
