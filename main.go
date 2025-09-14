package main

import (
	"fmt"
	"limitless-bot/commands"
	"limitless-bot/events"
	"limitless-bot/globals"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	_ "github.com/ncruces/go-sqlite3/vfs/memdb"
)

func main() {
	globals.Init()
	session, err := discordgo.New(fmt.Sprintf("Bot %s", globals.TOKEN))
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

	defer globals.DB.Close()
	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
