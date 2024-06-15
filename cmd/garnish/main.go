package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mocha8686/garnish/internal/model"
)

func main() {
	m := model.NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err);
		os.Exit(1)
	}
}
