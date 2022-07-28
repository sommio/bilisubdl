package bilibili

import (
	"errors"
	"fmt"

	"github.com/K0ng2/bilisubdl/utils"
)

// api url examples
/*
title
https://api.bilibili.tv/intl/gateway/web/v2/ogv/play/season_info?season_id=1049041

episode list
https://api.bilibili.tv/intl/gateway/web/v2/ogv/play/episodes?season_id=1049041

episode
https://api.bilibili.tv/intl/gateway/web/v2/subtitle?s_locale&episode_id=368729
*/

const bilibiliAPI string = "https://api.bilibili.tv/intl/gateway"
const bilibiliInfoAPI string = bilibiliAPI + "/web/v2/ogv/play/%s?season_id=%s"
// const bilibiliEpisodeAPI string = bilibiliAPI + "/subtitle?s_locale&episode_id="
const bilibiliEpisodeAPI string = bilibiliAPI + "/m/subtitle?ep_id="

func Info(id string) (*Info, error) {
	var info = new(Info)
	url := fmt.Sprintf(bilibiliInfoAPI, "season_info", id)
	if err := utils.ReqJson(info, url); err != nil {
		return nil, err
	}
	if info.Code != 0 {
		return nil, errors.New(fmt.Sprintf("Api response code %d: %s", info.Code, info.Message))
	}
	return info, nil
}

func Episodes(id string) (*Episodes, error) {
	var epList = new(Episodes)
	url := fmt.Sprintf(bilibiliInfoAPI, "episodes", id)
	if err := utils.ReqJson(epList, url); err != nil {
		return nil, err
	}
	if epList.Code != 0 {
		return nil, errors.New(fmt.Sprintf("Api response code %d: %s", epList.Code, epList.Message))
	}
	return epList, nil
}

func Episode(id string) (*Episode, error) {
	var ep = new(Episode)
	url := bilibiliEpisodeAPI + id
	if err := utils.ReqJson(ep, url); err != nil {
		return nil, err
	}
	if ep.Code != 0 {
		return nil, errors.New(fmt.Sprintf("Api response code %d: %s", ep.Code, ep.Message))
	}
	return ep, nil
}

func (s *Episode) Subtitle(language string) (string, error) {
	var index int
	var subJson = new(Subtitle)
	for i, s := range s.Data.Subtitles {
		if s.Key == lang {
			index = i
			break
		}
	}
	if index == 0 {
		return "", errors.New(fmt.Sprintf("Language \"%s\" not found", language))
	}
	if s.Data.Subtitles[index].IsMachine {
		log.Println("Warning machine translation")
	}
	err := utils.ReqJson(subJson, s.Data.Subtitles[index].URL)
	if err != nil {
		return "", err
	}
	return jsonToSRT(subJson), nil
}

// func SubToSRT(json *Subtitle) string {
// 	var sub string
// 	var content string
// 	for i, s := range json.Body {
// 		if s.Location == 2 {
// 			content = s.Content
// 		} else {
// 			content = fmt.Sprintf("{\\an%d}%s", s.Location, content)
// 		}
// 		sub += fmt.Sprintf("%d\n%s --> %s\n%s\n\n", i+1, utils.SecondToTime(s.From), utils.SecondToTime(s.To), content)
// 	}
// 	return sub
// }

func jsonToSRT(json BilibiliSubtitle) string {
	var sub string
	var content string
	for i, s := range json.Body {
		if i != 0 || i == len(json.Body) {
			sub += "\n\n"
		}
		content = s.Content
		if s.Location != 2 {
			content = fmt.Sprintf("{\\an%d}%s", s.Location, content)
		}
		sub += fmt.Sprintf("%d\n%s --> %s\n%s", i+1, utils.SecondToTime(s.From), utils.SecondToTime(s.To), content)
	}
	return sub
}
