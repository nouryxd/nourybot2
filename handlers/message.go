package handlers

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func PrivateMessage(message twitch.PrivateMessage, tc *twitch.Client) {
	if message.Message == "xd" {
		tc.Say(message.Channel, "xd")
	}
	fmt.Println(message)
}
