package modal

import (
	"github.com/bwmarrin/discordgo"
)

type TextInput struct {
	discordgo.TextInput
}

func newTextInput() *TextInput {
	return &TextInput{discordgo.TextInput{}}
}

func NewTextField(label string, id string, placeholder string, required bool) *TextInput {
	textInput := newTextInput().
		SetLabel(label).
		SetCustomID(id).
		SetPlaceholder(placeholder).
		SetRequired(required).
		SetStyle(discordgo.TextInputShort)

	return textInput
}

func NewTextArea(label string, id string, required bool) *TextInput {
	textInput := newTextInput().
		SetLabel(label).
		SetCustomID(id).
		SetRequired(required).
		SetStyle(discordgo.TextInputParagraph)

	return textInput
}

func (textInput *TextInput) SetLabel(label string) *TextInput {
	textInput.Label = label

	return textInput
}

func (textInput *TextInput) SetPlaceholder(placeholder string) *TextInput {
	textInput.Placeholder = placeholder

	return textInput
}

func (textInput *TextInput) SetStyle(style discordgo.TextInputStyle) *TextInput {
	textInput.Style = style

	return textInput
}

func (textInput *TextInput) SetCustomID(customID string) *TextInput {
	textInput.CustomID = customID

	return textInput
}

func (textInput *TextInput) SetRequired(required bool) *TextInput {
	textInput.Required = required

	return textInput
}
