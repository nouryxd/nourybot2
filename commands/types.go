package commands

// https://aws.random.cat/meow
type RandomCatResponse struct {
	File string `json:"file"`
}

// https://random.dog/woof.json
type RandomDogResponse struct {
	Url string `json:"url"`
}

// https://randomfox.ca/floof
type RandomFoxResponse struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

// https://xkcd.com/json.html
type XkcdResponse struct {
	Num       int    `json:"num"`
	SafeTitle string `json:"safe_title"`
	Img       string `json:"img"`
}

// https://api.twitch.tv/helix/streams/tags
type TwitchStreamTagsResponse struct {
	TagId                    string            `json:"tag_id"`
	IsAuto                   bool              `json:"is_auto"`
	LocalizationNames        map[string]string `json:"localization_names"`
	LocalizationDescriptions map[string]string `json:"localization_descriptions"`
}

type SubageResponse struct {
	User         string     `json:"user"`
	UserID       string     `json:"userid"`
	Channel      string     `json:"channel"`
	ChannelId    string     `json:"channelid"`
	SubageHidden bool       `json:"hidden"`
	Subscribed   bool       `json:"subscribed"`
	FollowedAt   string     `json:"followedAt"`
	Cumulative   Cumulative `json:"cumulative"`
	Streak       SubStreak  `json:"streak"`
}

type Cumulative struct {
	Months int `json:"months"`
}

type SubStreak struct {
	Months int `json:"months"`
}
