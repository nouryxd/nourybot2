package handlers

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

// PrivateMessage is called on each given incoming twitch.PrivateMessage
// and checks if the message contains something we should act on.
func PrivateMessage(message twitch.PrivateMessage, tc *twitch.Client) {
	if message.Message == "xd" {
		tc.Say(message.Channel, "xd")
	}
	fmt.Println(message)
}
