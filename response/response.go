package response

import "github.com/bwmarrin/discordgo"

type Response struct {
	ID                  string
	InteractionResponse *discordgo.InteractionResponse
}

func NewResponse() *Response {
	return &Response{InteractionResponse: &discordgo.InteractionResponse{}}
}

func NewInteractionResponse(id string, responseType discordgo.InteractionResponseType) *Response {
	response := NewResponse().
		SetApplicationCommandType(responseType)

	return response
}

func NewMessageResponse(id string) *Response {
	response := NewInteractionResponse(id, discordgo.InteractionResponseChannelMessageWithSource)

	return response
}

func NewUpdateMessageResponse(id string) *Response {
	response := NewInteractionResponse(id, discordgo.InteractionResponseUpdateMessage)

	return response
}

func (response *Response) SetID(id string) *Response {
	response.ID = id

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
