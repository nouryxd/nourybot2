package handlers

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func WhisperMessage(whisper twitch.WhisperMessage, client *twitch.Client) {
	if whisper.Message == "xd" {
		client.Whisper(whisper.User.Name, "xd")
	}
	fmt.Println(whisper)
}
