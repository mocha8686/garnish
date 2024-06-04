package context

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	color "github.com/fatih/color"
)

var prompt = color.BlueString("garnish") + "> "

var (
	errEmptyCommand   = errors.New("Empty command")
	errUnknownCommand = errors.New("Unknown command")
)

type Context struct {
	textInput textinput.Model
	history   []entry
}

type entry struct {
	cmd string
	out string
}

func NewContext() *Context {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 128
	ti.Prompt = prompt

	var hist []entry

	return &Context{
		textInput: ti,
		history:   hist,
	}
}

func (c *Context) Init() tea.Cmd {
	return textinput.Blink
}

func (c *Context) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			input := strings.TrimSpace(c.textInput.Value())
			c.textInput.SetValue("")

			res, err := c.handleCommand(input)
			if err == errEmptyCommand {
				c.appendHistory("", "")
				return c, nil
			} else if err == errUnknownCommand || errors.Unwrap(err) == errUnknownCommand {
				c.appendHistory(input, color.RedString(err.Error()))
				return c, nil
			} else if err != nil {
				log.Fatal(err)
			}

			c.appendHistory(input, res.out)
			return res.model, res.cmd
		case tea.KeyCtrlC, tea.KeyEsc:
			return c, tea.Quit
		}
	}

	c.textInput, cmd = c.textInput.Update(msg)
	return c, cmd
}

func (c *Context) View() string {
	if len(c.history) > 0 {
		return fmt.Sprintf("%v\n%v", c.formatHistory(), c.textInput.View())
	} else {
		return c.textInput.View()
	}
}

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
		fmt.Printf("Args: %v\n", args)
		return result{
			model: c,
			cmd:   nil,
			out:   "TODO: not yet implemented",
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

func (c *Context) appendHistory(cmd, out string) {
	e := entry{
		cmd: cmd,
		out: out,
	}
	c.history = append(c.history, e)
}

func (c *Context) formatHistory() string {
	var s strings.Builder
	l := len(c.history)

	for i, v := range c.history {
		s.WriteString(prompt)

		if len(v.cmd) > 0 || len(v.out) > 0 {
			s.WriteString(v.cmd + "\n" + v.out)
		}

		if i < l-1 {
			s.WriteRune('\n')
		}
	}
	return s.String()
}
