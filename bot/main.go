package bot

import (
	"fmt"

	twitch "github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-go/config"
	"github.com/lyx0/nourybot-go/handlers"
)

type Bot struct {
	cfg          *config.Config
	twitchClient *twitch.Client
	channels     map[string]*Channel
}

type Channel struct {
	Name     string
	Announce bool
}

var channels = map[string]*Channel{
	"nourybot": {Name: "nourybot", Announce: false},
	"nouryqt":  {Name: "nouryqt", Announce: false},
}

func NewBot(cfg *config.Config, twitchClient *twitch.Client) *Bot {
	return &Bot{
		cfg:          cfg,
		twitchClient: twitchClient,
		channels:     channels,
	}
}

// func (b *Bot) joinChannels() {
// 	for i := range channels {
// 		tc := b.twitchClient
// 		go tc.Connect()
// 		tc.Join(i)
// 		tc.Say(i, "xd")
// 		fmt.Println("Joined :", i)
// 	}
// }

// Create a new twitch client and return
// the connection to the caller
func (b *Bot) newClient() *twitch.Client {
	tc := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	return tc
}

func (b *Bot) Connect() {
	tc := b.newClient()

	// Connect to channels
	for i := range channels {
		tc.Join(i)
		tc.Say(i, "xd")
	}

	// Chat message has been received, forwarding it to
	// handlers.PrivateMessage
	tc.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println("xdasdsd")
		handlers.PrivateMessage(message)
	})

	tc.Connect()

}

// func (b *Bot) Connect() {
// 	tc := b.newClient()

// 	b.joinChannels()

// 	// Handle an incoming chat message
// 	tc.OnPrivateMessage(func(message twitch.PrivateMessage) {
// 		fmt.Println("xdasdsd")
// 		handlers.HandlePrivateMessage(message)
// 	})
// 	// Connect to the initial channels
// 	// for i := range channels {
// 	// 	tc.Join(i)
// 	// 	tc.Say(i, "xd")
// 	// }
// 	// Connnect to Twitch chat
// 	tc.Connect()
// }
