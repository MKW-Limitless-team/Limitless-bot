package button

import (
	em "limitless-bot/components/emoji"

	"github.com/bwmarrin/discordgo"
)

type Button struct {
	*discordgo.Button
}

func NewButton() *Button {
	return &Button{&discordgo.Button{}}
}

func NewBasicButton(label string, id string, style discordgo.ButtonStyle, disabled bool) *Button {
	button := NewButton().
		SetLabel(label).
		SetID(id).
		SetStyle(style).
		SetDisabled(disabled)

	return button
}

func NewEmojiButton(label string, id string, style discordgo.ButtonStyle, disabled bool, emoji string) *Button {
	button := NewButton().
		SetLabel(label).
		SetID(id).
		SetStyle(style).
		SetDisabled(disabled).
		SetEmoji(emoji)

	return button
}

func NewLinkButton(label string, url string, emoji string) *Button {
	button := NewButton().
		SetLabel(label).
		SetStyle(discordgo.LinkButton).
		SetURL(url).
		SetEmoji(emoji)

	return button
}

func (button *Button) SetLabel(label string) *Button {
	button.Label = label

	return button
}

func (button *Button) SetID(id string) *Button {
	button.CustomID = id

	return button
}

func (button *Button) SetStyle(style discordgo.ButtonStyle) *Button {
	button.Style = style

	return button
}

func (button *Button) SetDisabled(disabled bool) *Button {
	button.Disabled = disabled

	return button
}

func (button *Button) SetEmoji(emoji string) *Button {
	button.Emoji = em.NewBasicEmoji(emoji).ComponentEmoji

	return button
}

func (button *Button) SetURL(url string) *Button {
	button.URL = url

	return button
}
