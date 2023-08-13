package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eduwr/go_discord_bot/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}

	discordToken := os.Getenv("DISCORD_BOT_TOKEN")

	bot := bot.NewBot(discordToken)

	er := bot.Open()

	if er != nil {
		log.Fatal("Couldn't open a session with discord", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}
