package command

import "github.com/bwmarrin/discordgo"

type Option struct {
	*discordgo.ApplicationCommandOption
}

func newOption() *Option {
	return &Option{&discordgo.ApplicationCommandOption{}}
}

func NewCommandOption(name string, description string, optionType discordgo.ApplicationCommandOptionType, required bool) *Option {
	commandOption := newOption().
		SetName(name).
		SetDescription(description).
		SetOptionType(optionType).
		SetRequired(required)

	return commandOption
}

func (option *Option) AddChoice(choice *discordgo.ApplicationCommandOptionChoice) *Option {
	option.Choices = append(option.Choices, choice)

	return option
}

func (option *Option) SetAutoComplete(autoComplete bool) *Option {
	option.Autocomplete = autoComplete

	return option
}

func (option *Option) SetName(name string) *Option {
	option.Name = name

	return option
}

func (option *Option) SetDescription(description string) *Option {
	option.Description = description

	return option
}

func (option *Option) SetOptionType(optionType discordgo.ApplicationCommandOptionType) *Option {
	option.Type = optionType

	return option
}

func (option *Option) SetRequired(required bool) *Option {
	option.Required = required

	return option
}
