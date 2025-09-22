package main

import (
	"fmt"
	"limitless-bot/commands"
	"limitless-bot/events"
	"limitless-bot/globals"
	"limitless-bot/responses"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	_ "github.com/ncruces/go-sqlite3/vfs/memdb"
)

var (
	TOKEN = os.Getenv("BEPIS_TOKEN")
)

func main() {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", TOKEN))
	if err != nil {
		log.Fatal(err)
	}

	globals.Initialize(globals.SQLITEFILE)
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

	responses.RegisterResponses()

	defer session.Close()
	defer globals.GetConnection().Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
