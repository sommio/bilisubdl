package bilibili

import "time"

type Info struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Season struct {
			SeasonID       int    `json:"season_id"`
			Title          string `json:"title"`
			View           string `json:"view"`
			SeasonType     string `json:"season_type"`
			SeasonTypeEnum int    `json:"season_type_enum"`
			AliasName      string `json:"alias_name"`
			PlayerTime     string `json:"player_time"`
			OriginName     string `json:"origin_name"`
			IndexShow      string `json:"index_show"`
			Styles         []struct {
				ID    int    `json:"id"`
				Title string `json:"title"`
				Qs    string `json:"qs"`
			} `json:"styles"`
			UpdatePattern     string   `json:"update_pattern"`
			PayPolicy         string   `json:"pay_policy"`
			Description       string   `json:"description"`
			Actors            string   `json:"actors"`
			Staff             string   `json:"staff"`
			Directors         string   `json:"directors"`
			Writers           string   `json:"writers"`
			Performers        string   `json:"performers"`
			LimitAreas        []string `json:"limit_areas"`
			HorizontalCover   string   `json:"horizontal_cover"`
			AreaNames         string   `json:"area_names"`
			TotalEpisodesText string   `json:"total_episodes_text"`
			FirstEpisode      struct {
				EpisodeID    int    `json:"episode_id"`
				TitleDisplay string `json:"title_display"`
			} `json:"first_episode"`
			ViewHistory struct {
				EpisodeID    int    `json:"episode_id"`
				TitleDisplay string `json:"title_display"`
				Progress     int    `json:"progress"`
			} `json:"view_history"`
			IsFinished bool `json:"is_finished"`
		} `json:"season"`
	} `json:"data"`
}

type Episodes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Sections []struct {
			Title       string `json:"title"`
			EpListTitle string `json:"ep_list_title"`
			Episodes    []struct {
				Cover             string    `json:"cover"`
				Limit             int       `json:"limit"`
				LimitText         string    `json:"limit_text"`
				EpisodeID         int       `json:"episode_id"`
				ShortTitleDisplay string    `json:"short_title_display"`
				TitleDisplay      string    `json:"title_display"`
				PublishTime       time.Time `json:"publish_time"`
			} `json:"episodes"`
		} `json:"sections"`
	} `json:"data"`
}

type Episode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Subtitles []struct {
			URL        string `json:"url"`
			Lang       string `json:"lang"`
			LangKey    string `json:"lang_key"`
			SubtitleID int64  `json:"subtitle_id"`
		} `json:"subtitles"`
	} `json:"data"`
}

type Subtitle struct {
	FontSize        float64 `json:"font_size"`
	FontColor       string  `json:"font_color"`
	BackgroundAlpha float64 `json:"background_alpha"`
	BackgroundColor string  `json:"background_color"`
	Stroke          string  `json:"Stroke"`
	Body            []struct {
		From     float64 `json:"from"`
		To       float64 `json:"to"`
		Location int     `json:"location"`
		Content  string  `json:"content"`
	} `json:"body"`
}
