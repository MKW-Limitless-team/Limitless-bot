package container

import "github.com/bwmarrin/discordgo"

type TextDisplay struct {
	*discordgo.TextDisplay
}

func newTextDisplay() *TextDisplay {
	return &TextDisplay{&discordgo.TextDisplay{}}
}

func NewTextDisplay(content string) *TextDisplay {
	return newTextDisplay().SetContent(content)
}

func (textDisplay *TextDisplay) SetContent(content string) *TextDisplay {
	textDisplay.Content = content

	return textDisplay
}
