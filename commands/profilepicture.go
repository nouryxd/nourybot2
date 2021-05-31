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
type pfpApiResponse struct {
	Logo  string `json:"logo"`
	Error string `json:"error"`
}

func HandlePfp(channel string, username string) {
	resp, err := http.Get(fmt.Sprintf("https://api.ivr.fi/twitch/resolve/%s", username))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject pfpApiResponse
	json.Unmarshal(body, &responseObject)

	// User not found
	if responseObject.Error != "" {
		bot.SendTwitchMessage(channel, fmt.Sprintf(responseObject.Error+" FeelsBadMan"))
		return
	} else {
		bot.SendTwitchMessage(channel, fmt.Sprintf(responseObject.Logo))
	}
}
