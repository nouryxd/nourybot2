package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	bot "github.com/lyx0/nourybot-go/bot"
	cmd "github.com/lyx0/nourybot-go/commands"
	"github.com/lyx0/nourybot-go/db"
)

var channels = map[string]*bot.Channel{
	"nourybot": {Name: "nourybot"},
	// "nouryqt":              {Name: "nouryqt"},
	// "uudelleenkytkeytynyt": {Name: "uudelleenkytkeytynyt"},
	// "xnoury":               {Name: "xnoury"},
	// "nrybot":               {Name: "nrybot"},
	// "noemience":            {Name: "noemience"},
}

func connectToChannels() {
	for i := range channels {
		bot.Nourybot.Client.Join(i)
		log.Printf("Connected to channel: %v\n", i)
	}
	bot.SendTwitchMessage("nourybot", ":)")
	// bot.SendTwitchMessage("nouryqt", "pajaDink")
	// bot.SendTwitchMessage("nrybot", ":)")
	// bot.SendTwitchMessage("uudelleenkytkeytynyt", ":)")
	// bot.SendTwitchMessage("xnoury", "pajaDink")
}

func main() {
	log.Println("Starting")

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	botUser := os.Getenv("TWITCH_USER")
	botPass := os.Getenv("TWITCH_PASSWORD")

	bot.Nourybot = &bot.Bot{
		Client:   twitch.NewClient(botUser, botPass),
		Channels: channels,
		Uptime:   time.Now(),
	}

	bot.Nourybot.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// If channelID is missing something must have gone wrong.
		channelID := message.Tags["room-id"]
		if channelID == "" {
			fmt.Printf("Missing room-id tag in message")
			return
		}

		// So that the bot doesn't repeat itself.
		if message.Tags["user-id"] == "596581605" {
			return
		}

		// fmt.Printf("%v\n", message.Message)
		// fmt.Print(message.User.Badges)
		cmd.HandleMessage(message, bot.Nourybot)

	})

	// connectToChannels needs to be above err := Nourybot.Client.Connect()
	connectToChannels()

	// Connect to MySQL database
	db.Connect()

	// Connect to Twitch chat
	err := bot.Nourybot.Client.Connect()

	if err != nil {
		log.Fatalln(err.Error())
	}

}
