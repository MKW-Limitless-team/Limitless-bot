package response

import "github.com/bwmarrin/discordgo"

type Response struct {
	*discordgo.InteractionResponse
}

func newResponse() *Response {
	return &Response{InteractionResponse: &discordgo.InteractionResponse{}}
}

func newInteractionResponse(responseType discordgo.InteractionResponseType) *Response {
	response := newResponse().
		SetResponseType(responseType)

	return response
}

func NewAutoCompleteResponse() *Response {
	response := newInteractionResponse(discordgo.InteractionApplicationCommandAutocompleteResult)

	return response
}

func NewMessageResponse() *Response {
	response := newInteractionResponse(discordgo.InteractionResponseChannelMessageWithSource)

	return response
}

func NewModalResponse() *Response {
	response := newInteractionResponse(discordgo.InteractionResponseModal)

	return response
}

func NewUpdateMessageResponse() *Response {
	response := newInteractionResponse(discordgo.InteractionResponseUpdateMessage)

	return response
}

func (response *Response) SetResponseData(data *discordgo.InteractionResponseData) *Response {
	response.InteractionResponse.Data = data

	return response
}

func (response *Response) SetResponseType(responseType discordgo.InteractionResponseType) *Response {
	response.InteractionResponse.Type = discordgo.InteractionResponseType(responseType)

	return response
}
