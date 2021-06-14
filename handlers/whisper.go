package handlers

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func WhisperMessage(whisper twitch.WhisperMessage) {
	fmt.Println(whisper)
}
