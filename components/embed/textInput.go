package embed

import (
	"github.com/bwmarrin/discordgo"
)

type TextInput struct {
	*discordgo.TextInput
}

func newTextInput() *TextInput {
	return &TextInput{&discordgo.TextInput{}}
}

func NewTextField(label string, id string, placeholder string, required bool) *TextInput {
	return newTextInput().
		SetStyle(discordgo.TextInputShort)
}

func NewTextArea() *TextInput {
	return newTextInput().
		SetStyle(discordgo.TextInputParagraph)
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
