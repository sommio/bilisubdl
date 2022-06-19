package bilibili

import (
	"errors"
	"fmt"

	"github.com/K0ng2/bilisubdl/utils"
)

// api url examples
/*
title
https://api.bilibili.tv/intl/gateway/web/v2/ogv/play/season_info?season_id=1049041&platform=web

episode list
https://api.bilibili.tv/intl/gateway/web/v2/ogv/play/episodes?season_id=1049041&platform=web

episode
https://api.bilibili.tv/intl/gateway/web/v2/subtitle?&episode_id=368729
*/

const _API_URL string = "https://api.bilibili.tv/intl/gateway/web/v2"
const BASE_INFO_URL string = _API_URL + "/ogv/play/season_info?season_id=%s&platform=web"
const BASE_EPISODE_LIST_URL string = _API_URL + "/ogv/play/episodes?season_id=%s&platform=web"
const BASE_EPISODE_URL string = _API_URL + "/subtitle?&episode_id="

func GetInfo(id string) (*Info, error) {
	var info = new(Info)
	url := fmt.Sprintf(BASE_INFO_URL, id)
	if err := utils.GetJson(info, url); err != nil {
		return nil, err
	}

	if info.Data.Season.Title == "" {
		return nil, errors.New("Title not found")
	}

	return info, nil
}

func GetEpisodeList(id string) (*Episodes, error) {
	var epList = new(Episodes)
	url := fmt.Sprintf(BASE_EPISODE_LIST_URL, id)
	if err := utils.GetJson(epList, url); err != nil {
		return nil, err
	}
	return epList, nil
}

func GetEpisode(id int) (*Episode, error) {
	var ep = new(Episode)
	url := fmt.Sprintf("%s%d", BASE_EPISODE_URL, id)
	if err := utils.GetJson(ep, url); err != nil {
		return nil, err
	}
	return ep, nil
}

func (s *Episode) GetSubtitleJSON(language string) (*Subtitle, error) {
	var url string
	var subJson = new(Subtitle)
	for _, s := range s.Data.Subtitles {
		if s.LangKey == language {
			url = s.URL
		}
	}
	if url == "" {
		return nil, errors.New(fmt.Sprintf("Language \"%s\" not found", language))
	}
	err := utils.GetJson(subJson, url)
	if err != nil {
		return nil, err
	}
	return subJson, nil
}

func SubToSRT(json Subtitle) string {
	var sub string
	var content string
	for i, s := range json.Body {
		if s.Location == 2 {
			content = s.Content
		} else {
			content = fmt.Sprintf("{\\an%d}%s", s.Location, content)
		}
		sub += fmt.Sprintf("%d\n%s --> %s\n%s\n\n", i+1, utils.SecondToTime(s.From), utils.SecondToTime(s.To), content)
	}
	return sub
}
