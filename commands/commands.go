package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func HandleCommand(message twitch.PrivateMessage) {
	log.Info(message)
}
