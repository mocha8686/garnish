package context

import "strings"

type entry struct {
	cmd string
	out string
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
