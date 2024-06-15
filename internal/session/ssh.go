package session

import (
	"io"
	"os"

	"github.com/charmbracelet/x/term"
	"golang.org/x/crypto/ssh"
)

type Ssh struct {
	Addr         string
	Config       *ssh.ClientConfig
	conn         *ssh.Client
	shellSession *ssh.Session
}

func (s *Ssh) Connect() error {
	conn, err := ssh.Dial("tcp", s.Addr, s.Config)
	if err != nil {
		return err
	}
	s.conn = conn

	session, err := s.conn.NewSession()
	if err != nil {
		s.Close()
		return err
	}
	s.shellSession = session

	return nil
}

func (s *Ssh) Run() error {
	w, h, err := term.GetSize(os.Stdin.Fd())
	if err != nil {
		s.Close()
		return err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := s.shellSession.RequestPty("xterm", h, w, modes); err != nil {
		s.Close()
		return err
	}

	if err := s.shellSession.Shell(); err != nil {
		s.Close()
		return err
	}

	return s.shellSession.Wait()
}

func (s *Ssh) SetStdin(r io.Reader) {
	s.shellSession.Stdin = r
}

func (s *Ssh) SetStdout(w io.Writer) {
	s.shellSession.Stdout = w
}

func (s *Ssh) SetStderr(w io.Writer) {
	s.shellSession.Stderr = w
}

func (s *Ssh) Close() {
	s.shellSession.Close()
	s.shellSession = nil

	s.conn.Close()
	s.shellSession = nil
}
