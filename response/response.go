package response

import "github.com/bwmarrin/discordgo"

type Response struct {
	*discordgo.InteractionResponse
}

func newResponse() *Response {
	return &Response{InteractionResponse: &discordgo.InteractionResponse{}}
}

func NewInteractionResponse(responseType discordgo.InteractionResponseType) *Response {
	response := newResponse().
		SetApplicationCommandType(responseType)

	return response
}

func NewMessageResponse() *Response {
	response := NewInteractionResponse(discordgo.InteractionResponseChannelMessageWithSource)

	return response
}

func NewUpdateMessageResponse() *Response {
	response := NewInteractionResponse(discordgo.InteractionResponseUpdateMessage)

	return response
}

func (response *Response) SetInteractionResponseData(data *discordgo.InteractionResponseData) *Response {
	response.InteractionResponse.Data = data

	return response
}

func (response *Response) SetApplicationCommandType(responseType discordgo.InteractionResponseType) *Response {
	response.InteractionResponse.Type = discordgo.InteractionResponseType(responseType)

	return response
}
