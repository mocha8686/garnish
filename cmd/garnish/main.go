package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mocha8686/garnish/internal/context"
)

func main() {
	c := context.NewContext()
	p := tea.NewProgram(c)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err);
		os.Exit(1)
	}
}
