package response

import (
	embed "limitless-bot/components/embed"

	"github.com/bwmarrin/discordgo"
)

type Data struct {
	*discordgo.InteractionResponseData
}

func newData() *Data {
	return &Data{&discordgo.InteractionResponseData{}}
}

func NewResponseData(content string) *Data {
	responseData := newData().
		SetContent(content)

	return responseData
}

func NewFormData(title string, id string) *Data {
	responseData := newData().
		SetTitle(title).
		SetCustomID(id)

	return responseData
}

func NewAutoCompleteData(choices []*discordgo.ApplicationCommandOptionChoice) *Data {
	responseData := newData().
		SetChoices(choices)

	return responseData
}

func (data *Data) SetChoices(choices []*discordgo.ApplicationCommandOptionChoice) *Data {
	data.Choices = choices

	return data
}

func (data *Data) SetContent(content string) *Data {
	data.Content = content

	return data
}

func (data *Data) SetCustomID(id string) *Data {
	data.CustomID = id

	return data
}

func (data *Data) SetTitle(title string) *Data {
	data.Title = title

	return data
}

func (data *Data) AddEmbed(embed *embed.Embed) *Data {
	embeds := make([]*discordgo.MessageEmbed, 0)

	embeds = append(embeds, embed.MessageEmbed)

	data.Embeds = append(data.Embeds, embeds...)

	return data
}

func (data *Data) AddComponent(component discordgo.MessageComponent) *Data {
	data.Components = append(data.Components, component)

	return data
}
