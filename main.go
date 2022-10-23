package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// t := FromJSONFile("kanban.json")
	// fmt.Println(t)

	m := InitialModel()
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
