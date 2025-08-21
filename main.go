package main

import (
	"fmt"
	"limitless-bot/commands"
	"limitless-bot/events"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	TOKEN string
)

func main() {
	TOKEN = os.Getenv("SQUIRE_TOKEN")
	session, err := discordgo.New(fmt.Sprintf("Bot %s", TOKEN))
	if err != nil {
		log.Fatal(err)
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	events.RegisterEvents(session)

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = commands.RegisterCommands(session)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
