package commands

type RandomCatResponse struct {
	File string `json:"file"`
}

type RandomDogResponse struct {
	Url string `json:"url"`
}

type RandomFoxResponse struct {
	Image string `json:"url"`
	Link  string `json:"link"`
}
