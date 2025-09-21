package command

import "github.com/bwmarrin/discordgo"

type OptionChoice struct {
	*discordgo.ApplicationCommandOptionChoice
}

func newChoice() *OptionChoice {
	return &OptionChoice{&discordgo.ApplicationCommandOptionChoice{}}
}

func NewOptionChoice(name string, value any) *OptionChoice {
	optionChoice := newChoice().
		SetName(name).
		SetValue(value)

	return optionChoice
}

func (optionChoice *OptionChoice) SetName(name string) *OptionChoice {
	optionChoice.Name = name

	return optionChoice
}

func (optionChoice *OptionChoice) SetValue(value any) *OptionChoice {
	optionChoice.Value = value

	return optionChoice
}
