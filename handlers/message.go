package handlers

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func PrivateMessage(message twitch.PrivateMessage) {
	fmt.Println(message)
}
