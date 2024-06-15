package model

import (
	"errors"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/crypto/ssh"

	"github.com/mocha8686/garnish/internal/session"
)

var (
	errEmptyCommand   = errors.New("Empty command")
	errUnknownCommand = errors.New("Unknown command")
	errInvalidUsage   = errors.New("Invalid usage")
)

type result struct {
	model tea.Model
	cmd   tea.Cmd
}

func (m *Model) handleCommand(input string) (result, error) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return result{}, errEmptyCommand
	}

	cmd, args := fields[0], fields[1:]

	switch cmd {
	case "ssh":
		if args == nil || len(args) != 3 {
			return result{}, fmt.Errorf("%w: Usage: ssh <addr> <username> <password>", errInvalidUsage)
		}

		addr, username, password := args[0], args[1], args[2]
		// var hostKey ssh.PublicKey // TODO: implement

		s := session.Ssh{
			Addr: addr,
			Config: &ssh.ClientConfig{
				User:            username,
				HostKeyCallback: ssh.InsecureIgnoreHostKey(), // TODO: known_hosts
				Auth: []ssh.AuthMethod{
					ssh.Password(password),
				},
			},
		}

		if err := s.Connect(); err != nil {
			return result{}, err
		}

		return result{
			model: m,
			cmd:   tea.Exec(&s, nil),
		}, nil
	case "exit", "quit", "q":
		return result{
			model: m,
			cmd:   tea.Quit,
		}, nil
	default:
		return result{}, fmt.Errorf("%w: %v", errUnknownCommand, cmd)
	}
}
