package model

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

type Model struct {
	textInput textinput.Model
}

func NewModel() *Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 128
	ti.Prompt = prompt

	return &Model{
		textInput: ti,
	}
}

func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			input := strings.TrimSpace(m.textInput.Value())
			m.textInput.SetValue("")

			res, err := m.handleCommand(input)
			if err != nil {
				var cmd tea.Cmd

				if errors.Is(err, errEmptyCommand) {
					cmd = PromptPrintln()
				} else if errors.Is(err, errInvalidUsage) {
					cmd = PromptPrintf("%s\n%s", input, color.RedString(err.Error()))
				} else if errors.Is(err, errUnknownCommand) {
					cmd = PromptPrintf("%s\n%s", input, color.RedString(err.Error()))
				} else {
					log.Fatal(err)
				}

				return m, cmd
			}

			return res.model, tea.Batch(res.cmd, PromptPrintln(cmd))
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return m.textInput.View()
}

func PromptPrintf(template string, args ...any) tea.Cmd {
	return tea.Printf(prompt+template, args...)
}

func PromptPrintln(args ...any) tea.Cmd {
	return tea.Println(prompt + fmt.Sprint(args...))
}
