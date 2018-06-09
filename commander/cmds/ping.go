package cmds

import (
	"../../commander"
)

type Ping struct{}

func (cmd *Ping) Execute(c *commander.Context, args []string) {
	c.Reply("Pong!")
}
func (cmd *Ping) Name() string { return "ping" }
