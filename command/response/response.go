package response

import "github.com/bwmarrin/discordgo"

type Response struct {
	*discordgo.InteractionResponse
}

func NewResponse() *Response {
	return &Response{&discordgo.InteractionResponse{}}
}

func NewInteractionResponse(responseType discordgo.InteractionResponseType) *Response {
	response := NewResponse().
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
	response.Data = data

	return response
}

func (response *Response) SetApplicationCommandType(responseType discordgo.InteractionResponseType) *Response {
	response.Type = discordgo.InteractionResponseType(responseType)

	return response
}
