package main

import (
	"fmt"
	"limitless-bot/commands"
	"limitless-bot/events"
	"limitless-bot/globals"
	"limitless-bot/responses"
	"limitless-bot/utils"
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
	TOKEN      = os.Getenv("BEPIS_TOKEN")
	ADMIN_ROLE = os.Getenv("ADMIN_ROLE")
	SECRET     = os.Getenv("SECRET")
)

func main() {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", TOKEN))
	if err != nil {
		log.Fatal(err.Error())
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	if ADMIN_ROLE == "" {
		log.Fatal("No admin role set")
	}

	globals.ADMIN_ROLE = ADMIN_ROLE

	if ADMIN_ROLE == "" {
		log.Fatal("No secret set")
	}
	globals.SECRET = SECRET

	events.RegisterEvents(session)

	err = session.Open()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = commands.RegisterCommands(session)
	if err != nil {
		log.Fatal(err.Error())
	}

	responses.RegisterResponses()

	utils.Modes = utils.PopulateRandomOptions("./events.csv", utils.Modes)
	utils.Modifiers = utils.PopulateRandomOptions("./modifiers.csv", utils.Modifiers)
	utils.Tracks = utils.PopulateRandomOptions("./tracks.csv", utils.Tracks)
	utils.FolderNames = utils.PopulateFolderNames("./folderNames.csv", utils.FolderNames)

	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
