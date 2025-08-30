package modal

import (
	"limitless-bot/components/emoji"

	"github.com/bwmarrin/discordgo"
)

type SelectMenuOption struct {
	discordgo.SelectMenuOption
}

func newMenuOption() SelectMenuOption {
	return SelectMenuOption{discordgo.SelectMenuOption{}}
}

func NewSelectMenuOption(label string, value string) SelectMenuOption {
	selectOption := newMenuOption().
		SetLabel(label).
		SetValue(value)

	return selectOption
}

func NewEmojiMenuOption(label string, value string, description string, emoji *emoji.Emoji) SelectMenuOption {
	selectOption := newMenuOption().
		SetLabel(label).
		SetValue(value).
		SetDescription(description).
		SetEmoji(emoji)

	return selectOption
}

func (selectOption SelectMenuOption) SetLabel(label string) SelectMenuOption {
	selectOption.Label = label

	return selectOption
}

func (selectOption SelectMenuOption) SetValue(value string) SelectMenuOption {
	selectOption.Value = value

	return selectOption
}

func (selectOption SelectMenuOption) SetDescription(description string) SelectMenuOption {
	selectOption.Description = description

	return selectOption
}

func (selectOption SelectMenuOption) SetEmoji(emoji *emoji.Emoji) SelectMenuOption {
	selectOption.Emoji = emoji.ComponentEmoji

	return selectOption
}
