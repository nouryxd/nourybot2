package bot

import (
	"database/sql"
	"fmt"
	"log"

	twitch "github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot-go/config"
	"github.com/lyx0/nourybot-go/handlers"
)

type Bot struct {
	cfg          *config.Config
	twitchClient *twitch.Client
	sqlClient    *sql.DB
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

// Newbot returns a pointer to a Bot from a given
// *config.Config and *twitch.Client
func NewBot(cfg *config.Config, twitchClient *twitch.Client, sqlClient *sql.DB) *Bot {
	return &Bot{
		cfg:          cfg,
		twitchClient: twitchClient,
		sqlClient:    sqlClient,
		channels:     channels,
	}
}

// newClient creates a new client from a  given *twitch.Client
func (b *Bot) newClient() *twitch.Client {
	tc := twitch.NewClient(b.cfg.Username, b.cfg.Oauth)
	return tc
}

// Connect connects to chat and listen for incoming messages
func (b *Bot) Connect() error {
	tc := b.newClient()

	// Connect to channels
	for i := range channels {
		tc.Join(i)
		tc.Say(i, "xd")
		fmt.Printf("Connected to: #%s\n", i)
	}

	// OnPrivateMessage forwards the received twitch.PrivateMessage
	// to the appropiate PrivateMessage handler function.
	tc.OnPrivateMessage(func(message twitch.PrivateMessage) {
		handlers.PrivateMessage(message, tc)
	})

	// OnWhisperMessage forwards the received twitch.WhisperMessage
	// to the appropiate WhisperMessage handler function.
	tc.OnWhisperMessage(func(whisper twitch.WhisperMessage) {
		handlers.WhisperMessage(whisper, tc)
	})

	// Actually connect to chat and return the connection
	err := tc.Connect()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
