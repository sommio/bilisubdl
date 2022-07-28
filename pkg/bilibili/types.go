package bilibili

import "encoding/json"

type Info struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Season struct {
			Title string `json:"title"`
		} `json:"season"`
	} `json:"data"`
}

type Episodes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Sections []struct {
			EpListTitle string `json:"ep_list_title"`
			Episodes    []struct {
				// ShortTitleDisplay string      `json:"short_title_display"`
				EpisodeID    json.Number `json:"episode_id"`
				TitleDisplay string      `json:"title_display"`
			} `json:"episodes"`
		} `json:"sections"`
	} `json:"data"`
}

type Episode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Subtitles []struct {
			URL     	string `json:"url"`
			// Lang    string `json:"lang"`
			Title     string `json:"title"`
			// LangKey string `json:"lang_key"`
			Key       string `json:"key"`
			IsMachine bool   `json:"is_machine"`
		} `json:"subtitles"`
	} `json:"data"`
}

type Subtitle struct {
	Body    []struct {
		From     float64 `json:"from"`
		To       float64 `json:"to"`
		Location int     `json:"location"`
		Content  string  `json:"content"`
	} `json:"body"`
}
