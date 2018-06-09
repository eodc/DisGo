package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	discord, err := discordgo.New("Bot NDU0Nzg3OTE3MDcyNjk1Mjk4.DfyhyQ.y9j_8xa1jlpGrp2ekjQr9zFaFIM")
	if err != nil {
		fmt.Println(err)
		return
	}
	discord.AddHandler(handleMessage)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("LINK START")

	discord.UpdateStatus(0, "Go!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("\nGOODBYE")
	discord.Close()
}

func handleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID ||
		!strings.HasPrefix(message.Content, "!") {
		return
	}
	command := strings.ToLower(strings.TrimPrefix(message.Content, "!"))
	if command == "ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}
