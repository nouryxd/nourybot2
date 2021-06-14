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

func (b *Bot) newClient() *twitch.Client {
	tc := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	return tc
}

func (b *Bot) Connect() {
	tc := b.newClient()

	// Connect to the initial channels
	tc.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println("xdasdsd")
		handlers.HandlePrivateMessage(message)
	})
	for i := range channels {
		tc.Join(i)
		tc.Say(i, "xd")
	}
	// b.connectToChannels()
	tc.Connect()
}

// func (b *Bot) connectToChannels() {
// 	tc := b.twitchClient
// 	for i := range channels {
// 		tc.Join(i)
// 		tc.Say(i, "xd")
// 		fmt.Println("Connected to:", i)
// 	}
// }
