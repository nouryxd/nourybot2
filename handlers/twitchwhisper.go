package handlers

import (
	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func TwitchWhisper(whisper twitch.WhisperMessage, twitchClient *twitch.Client) {
	log.Info(whisper)
	if whisper.Message == "xd" {
		twitchClient.Whisper(whisper.User.Name, "xd")
	}

}
