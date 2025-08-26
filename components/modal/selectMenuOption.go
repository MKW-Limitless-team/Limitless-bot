package modal

import "github.com/bwmarrin/discordgo"

type SelectMenuOption struct {
	discordgo.SelectMenuOption
}

func newMenuOption() *SelectMenuOption {
	return &SelectMenuOption{discordgo.SelectMenuOption{}}
}

func NewSelectMenuOption() {

}
