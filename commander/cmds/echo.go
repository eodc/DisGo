package cmds

import "../../commander"

type Echo struct{}

func (cmd *Echo) Execute(c *commander.Context, args []string) {
	payload := ""
	for _, value := range args {
		payload += value + " "
	}
	if payload != "" {
		c.SendMessage(payload)
	}
}

func (cmd *Echo) Name() string { return "echo" }
