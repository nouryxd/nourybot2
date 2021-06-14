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

func (b *Bot) joinChannels() {
	for i := range channels {
		tc := b.newClient()
		go tc.Connect()
		tc.Join(i)
		tc.Say(i, "xd")
		fmt.Println("Joined :", i)
	}
}

// Create a new twitch client and return
// the connection to the caller
func (b *Bot) newClient() *twitch.Client {
	tc := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	return tc
}

func (b *Bot) Connect() {
	tc := b.newClient()

	// Handle an incoming chat message
	tc.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println("xdasdsd")
		handlers.HandlePrivateMessage(message)
	})
	b.joinChannels()

	// Connect to the initial channels
	// for i := range channels {
	// 	tc.Join(i)
	// 	tc.Say(i, "xd")
	// }
	// Connnect to Twitch chat
	tc.Connect()
}
