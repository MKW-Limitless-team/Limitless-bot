package embed

import (
	"limitless-bot/components/button"

	"github.com/bwmarrin/discordgo"
)

type ActionRow struct {
	*discordgo.ActionsRow
}

func NewActionRow() *ActionRow {
	return &ActionRow{&discordgo.ActionsRow{}}
}

func (actionRow *ActionRow) AddButton(button *button.Button) *ActionRow {
	actionRow.Components = append(actionRow.Components, button)

	return actionRow
}
