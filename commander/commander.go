package commander

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Execute(*Context)
	Name() string
}

type Context struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
}

func (c *Context) Reply(content string) {
	c.Session.ChannelMessageSend(c.Message.ChannelID, fmt.Sprintf("%s, %s",
		c.Message.Author.Mention(), content))
}

var commands map[string]Command

func AddCommand(c Command) {
	if commands == nil {
		commands = make(map[string]Command)
	}
	commands[c.Name()] = c
}

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	payload := strings.ToLower(strings.TrimPrefix(m.Content, "!"))
	if commands != nil {
		for name, cmd := range commands {
			if name == payload {
				cmd.Execute(&Context{s, m})
				return
			}
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unknown Command: `%s`", payload))
	}
}
