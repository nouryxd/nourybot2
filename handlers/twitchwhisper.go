package handlers

import (
	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

// WhisperMessage is called on every given incoming twitch.WhisperMessage
// and checks if the whisper contains something we should act on.
func WhisperMessage(whisper twitch.WhisperMessage, tc *twitch.Client) {
	if whisper.Message == "xd" {
		tc.Whisper(whisper.User.Name, "xd")
	}
	log.Info(whisper)
}
