package client

import (
	"database/sql"

	twitch "github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot2/config"
	"github.com/lyx0/nourybot2/db"
	"github.com/lyx0/nourybot2/handlers"
	log "github.com/sirupsen/logrus"
)

type TwitchBot struct {
	cfg          *config.Config
	twitchClient *twitch.Client
	sqlClient    *sql.DB
}

type Channel struct {
	Name     string
	Announce bool
}

// Newbot returns a pointer to a Bot from a given
// *config.Config and *twitch.Client.
func NewTwitchBot(cfg *config.Config, twitchClient *twitch.Client, sqlClient *sql.DB) *TwitchBot {
	return &TwitchBot{
		cfg:          cfg,
		twitchClient: twitchClient,
		sqlClient:    sqlClient,
	}
}

// newClient creates a new client from a  given *twitch.Client.
func (tb *TwitchBot) newTwitchClient() *twitch.Client {
	tc := twitch.NewClient(tb.cfg.Username, tb.cfg.Oauth)
	return tc
}

// Connect connects to chat and listen for incoming messages.
func (tb *TwitchBot) Connect() error {
	tc := tb.newTwitchClient()

	// Get a list of channels from the database to join.
	db.JoinChannels(tc, tb.sqlClient)

	// Get a list of channels from the database in which
	// we should announce when we join.
	db.AnnounceJoin(tc, tb.sqlClient)

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

	// Actually connect to chat and return the connection.
	err := tc.Connect()
	if err != nil {
		log.Fatal(err)
		// log.Fatal(err)
		return err
	}

	return err
}
