package commander

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/shlex"
)

type Command interface {
	Execute(*Context, []string)
	Name() string
}

type Context struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
}

func (c *Context) Reply(content string) {
	c.SendMessage(fmt.Sprintf("%s, %s", c.Message.Author.Mention(), content))
}

func (c *Context) SendMessage(content string) {
	c.Session.ChannelMessageSend(c.Message.ChannelID, content)
}

var commands map[string]Command

func AddCommand(c Command) {
	if commands == nil {
		commands = make(map[string]Command)
	}
	commands[c.Name()] = c
}

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	fullCmd, err := shlex.Split(strings.ToLower(strings.TrimPrefix(m.Content, "!")))

	if err != nil {
		fmt.Println("[ERR] %v", err)
		return
	}

	if len(fullCmd) > 0 {
		name := fullCmd[0]
		var args []string

		if len(fullCmd) > 1 {
			args = fullCmd[1:]
		}

		cmd := commands[name]
		if cmd != nil {
			cmd.Execute(&Context{s, m}, args)
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, `%s` is an invalid command.",
				m.Author.Mention(), name))
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, you didn't send a command!", m.Author.Mention()))
	}

}
