package bot

import (
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

type Bot struct {
	Client *twitch.Client
	// Channels map[string]*Channel
	UserID string
	Uptime time.Time
}

type Channel struct {
	Name    string
	LastMsg string
}

type Command struct {
	Name string
	Run  func(msg twitch.PrivateMessage, args []string)
}

var (
	Nourybot *Bot
)

// SendTwitchMessage sends a twitch message to a given
// target(channel) string and message string
func SendTwitchMessage(target string, message string) {
	if len(message) == 0 {
		return
	}

	// if message[0] == '.' || message[0] == '/' {
	// 	message = ". " + message
	// }

	if len(message) > 247 {
		firstMessage := message[0:499]
		secondMessage := message[499:]
		Nourybot.Client.Say(target, firstMessage)
		Nourybot.Client.Say(target, secondMessage)
		return
	}

	Nourybot.Client.Say(target, message)
}
