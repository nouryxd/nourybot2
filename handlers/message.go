package handlers

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func HandlePrivateMessage(message twitch.PrivateMessage) {
	fmt.Println(message)
}
