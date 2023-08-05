package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}

	discordToken := os.Getenv("DISCORD_BOT_TOKEN")

	discord, err := discordgo.New("Bot " + discordToken)

	if err != nil {
		log.Fatal("Couldn't create a session with discord", err)
	}

	discord.Identify.Intents = discordgo.IntentMessageContent
	// handle the incomming messages
	discord.AddHandler(messageCreate)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		log.Fatal("Couldn't create a websocket connection to Discord")
	}

	fmt.Println("Bot is live")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("MESSAGE RECEIVED", m.Content)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!gopher" {
		s.ChannelMessageSend(m.ChannelID, "Hello Go Bot")
	}
}
