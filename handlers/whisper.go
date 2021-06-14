package handlers

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

// WhisperMessage is called on every given incoming twitch.WhisperMessage
// and checks if the whisper contains something we should act on.
func WhisperMessage(whisper twitch.WhisperMessage, tc *twitch.Client) {
	if whisper.Message == "xd" {
		tc.Whisper(whisper.User.Name, "xd")
	}
	fmt.Println(whisper)
}
