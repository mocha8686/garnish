package context

import (
	"fmt"
	"log"
	"strings"

	color "github.com/fatih/color"
	repl "github.com/openengineer/go-repl"
)

type Context struct {
	r *repl.Repl
}

func NewContext() *Context {
	c := &Context{}
	c.r = repl.NewRepl(c)
	return c
}

func (c *Context) Start() {
	if err := c.r.Loop(); err != nil {
		log.Fatal(err)
	}
}

func (c *Context) Prompt() string {
	return color.BlueString("garnish") + "> "
}

func (c *Context) Tab(buffer string) string {
	// TODO: implement
	return ""
}

func (c *Context) Eval(line string) string {
	fields := strings.Fields(line)

	if len(fields) == 0 {
		return ""
	}

	cmd, args := fields[0], fields[1:]

	switch cmd {
	case "ssh":
		return fmt.Sprintf("TODO: not yet implemented | args: %s", args)
	case "exit", "quit":
		c.r.Quit()
		return ""
	default:
		return color.RedString(fmt.Sprintf("Unrecognized command `%s`.", cmd))
	}
}
