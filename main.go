package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot2/bot"
	"github.com/lyx0/nourybot2/config"
	"github.com/lyx0/nourybot2/db"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func main() {
	cfg := config.LoadConfig()

	sqlClient := db.Connect(cfg)
	defer sqlClient.Close()

	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)

	bot := bot.NewBot(cfg, twitchClient, sqlClient)

	wg.Add(2)

	go func() {
		log.Info("Connecting to Twitch")

		err := bot.ConnectTwitch()
		if err != nil {
			log.Fatal("Couldn't connect to Twitch", err)
			os.Exit(1)
		}
		wg.Done()
	}()


	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-c

	sig := <-c
	log.Info("Got signal:", sig)
	bot.CloseConnection()
	os.Exit(0)
}
