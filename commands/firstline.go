package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bot "github.com/lyx0/nourybot-go/bot"
)

// https://api.ivr.fi
type flApiResponse struct {
	User    string `json:"user"`
	Message string `json:"message"`
	Time    string `json:"time"`
	Error   string `json:"error"`
}

func HandleFirstline(channel string, username string, streamer string) {
	resp, err := http.Get(fmt.Sprintf("https://api.ivr.fi/logs/firstmessage/%s/%s", streamer, username))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject flApiResponse
	json.Unmarshal(body, &responseObject)

	// User or channel was not found
	if responseObject.Error != "" {
		bot.SendTwitchMessage(channel, fmt.Sprintf(responseObject.Error+" FeelsBadMan"))
		return
	} else {
		bot.SendTwitchMessage(channel, fmt.Sprintf(username+": "+responseObject.Message+" ("+responseObject.Time+" ago)."))
	}
}
