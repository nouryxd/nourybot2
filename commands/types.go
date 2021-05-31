package commands

// https://api.twitch.tv/helix/streams/tags
type TwitchStreamTagsResponse struct {
	TagId                    string            `json:"tag_id"`
	IsAuto                   bool              `json:"is_auto"`
	LocalizationNames        map[string]string `json:"localization_names"`
	LocalizationDescriptions map[string]string `json:"localization_descriptions"`
}
