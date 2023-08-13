package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/eduwr/go_discord_bot/rollingboard"
)

type Bot struct {
	session *discordgo.Session
}

func NewBot(token string) *Bot {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Couldn't create a session with Discord", err)
	}

	bot := &Bot{
		session: session,
	}

	session.Identify.Intents = discordgo.IntentMessageContent
	session.Identify.Intents = discordgo.IntentsGuildMessages
	session.AddHandler(bot.messageCreate)

	return bot
}

func (bot *Bot) Open() error {
	return bot.session.Open()
}

func (bot *Bot) Close() {
	bot.session.Close()
}

func (bot *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("MESSAGE RECEIVED", m.Content)
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!gopher" {
		s.ChannelMessageSend(m.ChannelID, "Hello Go Bot")
	}

	if strings.HasPrefix(m.Content, "!roll") {
		board, err := rollingboard.NewRollingBoard(m.Content)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("error: %v", err))

		}

		r := board.RollDices()

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s => %d", m.Content, r))

	}
}
