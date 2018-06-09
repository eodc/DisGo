package commander

type Ping struct{}

func (cmd *Ping) Execute(c *Context) {
	c.Reply("Pong!")
}
func (cmd *Ping) Name() string { return "ping" }
