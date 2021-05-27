package commands

type RandomCatResponse struct {
	File string `json:"file"`
}

type RandomDogResponse struct {
	Url string `json:"url"`
}

type RandomFoxResponse struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

type XkcdResponse struct {
	Num       int    `json:"num"`
	SafeTitle string `json:"safe_title"`
	Img       string `json:"img"`
}
