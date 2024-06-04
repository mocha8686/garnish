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

const name = "garnish"

var prompt = color.BlueString(name) + "> "

type Context struct {
	textInput textinput.Model
	history   []entry
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
			if errors.Is(err, errEmptyCommand) {
				c.appendHistory("", "")
				return c, nil
			} else if errors.Is(err, errUnknownCommand) {
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
