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
type subageApiResponse struct {
	User         string     `json:"user"`
	UserID       string     `json:"userid"`
	Channel      string     `json:"channel"`
	ChannelId    string     `json:"channelid"`
	SubageHidden bool       `json:"hidden"`
	Subscribed   bool       `json:"subscribed"`
	FollowedAt   string     `json:"followedAt"`
	Cumulative   Cumulative `json:"cumulative"`
	Streak       SubStreak  `json:"streak"`
	Error        string     `json:"error"`
}

type Cumulative struct {
	Months int `json:"months"`
}

type SubStreak struct {
	Months int `json:"months"`
}

func HandleSubage(channel string, username string, streamer string) {
	resp, err := http.Get(fmt.Sprintf("https://api.ivr.fi/twitch/subage/%s/%s", username, streamer))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var responseObject subageApiResponse
	json.Unmarshal(body, &responseObject)

	// User or channel was not found
	if responseObject.Error != "" {
		bot.SendTwitchMessage(channel, fmt.Sprintf(responseObject.Error+" FeelsBadMan"))
		return
	}
	// User was found but has their subscription hidden.
	if responseObject.SubageHidden {
		bot.SendTwitchMessage(channel, fmt.Sprintf(username+" has their subscription status hidden. FeelsBadMan"))
	} else {
		months := fmt.Sprint(responseObject.Cumulative.Months)
		bot.SendTwitchMessage(channel, fmt.Sprintf(username+" has been subscribed to "+streamer+" for "+months+" months."))
	}
}
