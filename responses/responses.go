package responses

import (
	"limitless-bot/response"
)

var (
	CommandResponses   []*response.Response = make([]*response.Response, 0)
	ComponentResponses []*response.Response = make([]*response.Response, 0)
)

func RegisterResponses() {

	// Add command responses here
	// CommandResponses = append(CommandResponses, HelpResponse())
}
