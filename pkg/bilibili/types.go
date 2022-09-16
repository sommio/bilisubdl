package bilibili

import (
	"encoding/json"
	"time"
)

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
				ShortTitleDisplay string      `json:"short_title_display"`
				EpisodeID         json.Number `json:"episode_id"`
				TitleDisplay      string      `json:"title_display"`
				PublishTime       time.Time   `json:"publish_time"`
			} `json:"episodes"`
		} `json:"sections"`
	} `json:"data"`
}

type Episode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Subtitles []struct {
			URL string `json:"url"`
			// Lang    string `json:"lang"`
			Title string `json:"title"`
			// LangKey string `json:"lang_key"`
			Key       string `json:"key"`
			IsMachine bool   `json:"is_machine"`
		} `json:"subtitles"`
	} `json:"data"`
}

type Subtitle struct {
	Body []struct {
		From     float64 `json:"from"`
		To       float64 `json:"to"`
		Location int     `json:"location"`
		Content  string  `json:"content"`
	} `json:"body"`
}

type Timeline struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// TTL     int    `json:"ttl"`
	Data struct {
		Items []struct {
			DayOfWeek string `json:"day_of_week"`
			IsToday   bool   `json:"is_today"`
			// DateText     string `json:"date_text"`
			FullDateText string `json:"full_date_text"`
			Cards        []struct {
				// Type        string      `json:"type"`
				// CardType    string      `json:"card_type"`
				Title string `json:"title"`
				// Cover       string      `json:"cover"`
				// View        string      `json:"view"`
				// Styles      string      `json:"styles"`
				// StyleList   interface{} `json:"style_list"`
				SeasonID string `json:"season_id"`
				// EpisodeID   string      `json:"episode_id"`
				IndexShow string `json:"index_show"`
				// Label       int         `json:"label"`
				// RankInfo    interface{} `json:"rank_info"`
				// ViewHistory interface{} `json:"view_history"`
				// Watched     string      `json:"watched"`
				// Duration    string      `json:"duration"`
				// ViewAt      string      `json:"view_at"`
				// PubTimeText string `json:"pub_time_text"`
				// Unavailable bool        `json:"unavailable"`
			} `json:"cards"`
		} `json:"items"`
	} `json:"data"`
}

type Search struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// TTL     int    `json:"ttl"`
	Data []struct {
		Module string `json:"module"`
		// Title   string `json:"title"`
		// HasNext bool   `json:"has_next"`
		// TrackID string `json:"track_id"`
		Items []struct {
			Title    string      `json:"title"`
			SeasonID json.Number `json:"season_id"`
			// Mid            int         `json:"mid"`
			// Desc           string      `json:"desc"`
			// Followers      string      `json:"followers"`
			// ArchiveNum     string      `json:"archive_num"`
			// Duration       string      `json:"duration"`
			// Cover          string      `json:"cover"`
			// UpFollowButton interface{} `json:"up_follow_button"`
			// Score          int         `json:"score"`
			// TitleMatch     interface{} `json:"title_match"`
			// TabType        int         `json:"tab_type"`
			// TabName        string      `json:"tab_name"`
			// Aid            int         `json:"aid"`
			// Avatar         string      `json:"avatar"`
			// Label          int         `json:"label"`
			IndexShow string `json:"index_show"`
			// View           string      `json:"view"`
			// Styles         string      `json:"styles"`
		} `json:"items"`
	} `json:"data"`
}
