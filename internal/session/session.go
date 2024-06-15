package session

import (
	"io"

	tea "github.com/charmbracelet/bubbletea"
)

type Session interface {
	Connect() error
	tea.ExecCommand
	io.Closer
}
