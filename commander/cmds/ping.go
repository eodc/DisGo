package cmds

import (
	"../../commander"
)

type Ping struct{}

func (cmd *Ping) Execute(c *commander.Context) {
	c.Reply("Pong!")
}
func (cmd *Ping) Name() string { return "ping" }
