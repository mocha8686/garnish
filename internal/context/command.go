package context

import (
	"errors"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	errEmptyCommand   = errors.New("Empty command")
	errUnknownCommand = errors.New("Unknown command")
)

type result struct {
	model tea.Model
	cmd   tea.Cmd
	out   string
}

func (c *Context) handleCommand(input string) (result, error) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return result{}, errEmptyCommand
	}

	cmd, args := fields[0], fields[1:]

	switch cmd {
	case "ssh":
		var output strings.Builder

		output.WriteString(fmt.Sprintf("Args: %v\n", strings.Join(args, ", ")))
		output.WriteString("TODO: not yet implemented")

		return result{
			model: c,
			cmd:   nil,
			out:   output.String(),
		}, nil
	case "exit", "quit", "q":
		return result{
			model: c,
			cmd:   tea.Quit,
			out:   "Goodbye.",
		}, nil
	default:
		return result{}, fmt.Errorf("%w: %v", errUnknownCommand, cmd)
	}
}
