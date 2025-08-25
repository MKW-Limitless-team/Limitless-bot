package components

import (
	"github.com/bwmarrin/discordgo"
)

type ActionRow struct {
	*discordgo.ActionsRow
}

func NewActionRow() *ActionRow {
	return &ActionRow{&discordgo.ActionsRow{}}
}

func (actionRow *ActionRow) AddComponent(component discordgo.MessageComponent) *ActionRow {
	actionRow.Components = append(actionRow.Components, component)

	return actionRow
}
