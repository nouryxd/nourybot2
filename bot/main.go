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
}

func NewBot(cfg *config.Config, twitchClient *twitch.Client) *Bot {
	return &Bot{
		cfg:          cfg,
		twitchClient: twitchClient,
	}
	// err := twitchClient.Connect()
	// if err != nil {
	// 	fmt.Printf("Error connecting to Twitch chat %#v", err)
	// }
}

func (b *Bot) Connect(channel string) {
	tc := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	tc.Join(channel)
	tc.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println("xdasdsd")
		handlers.HandlePrivateMessage(message)
	})
	tc.Connect()
}
