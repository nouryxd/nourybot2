package handlers

import (
	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func TwitchMessage(message twitch.PrivateMessage, tc *twitch.Client) {
	log.Info(message)
	if message.Message == "xd" {
		tc.Say(message.Channel, "xd")
	}
}
