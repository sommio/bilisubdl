package bilibili

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

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

func Info(id string) (*BilibiliInfo, error) {
	var info = new(BilibiliInfo)
	url := fmt.Sprintf(bilibiliInfoAPI, "season_info", id)
	resp, err := utils.Request(url)
	if err != nil {
		return nil, err
	}
	if resp.Json(info); err != nil {
		return nil, err
	}
	if info.Code != 0 {
		return nil, errors.New(fmt.Sprintf("Api response code %d: %s", info.Code, info.Message))
	}
	return info, nil
}

func Episodes(id string) (*BilibiliEpisodes, error) {
	var epList = new(BilibiliEpisodes)
	url := fmt.Sprintf(bilibiliInfoAPI, "episodes", id)
	resp, err := utils.Request(url)
	if err != nil {
		return nil, err
	}
	if resp.Json(epList); err != nil {
		return nil, err
	}
	if epList.Code != 0 {
		return nil, errors.New(fmt.Sprintf("Api response code %d: %s", epList.Code, epList.Message))
	}
	return epList, nil
}

func Episode(id string) (*BilibiliEpisode, error) {
	var ep = new(BilibiliEpisode)
	url := bilibiliEpisodeAPI + id
	resp, err := utils.Request(url)
	if err != nil {
		return nil, err
	}
	if resp.Json(ep); err != nil {
		return nil, err
	}
	if ep.Code != 0 {
		return nil, errors.New(fmt.Sprintf("Api response code %d: %s", ep.Code, ep.Message))
	}
	return ep, nil
}

func (s *BilibiliEpisode) Subtitle(language string) ([]byte, string, error) {
	var index int = -1
	for i, s := range s.Data.Subtitles {
		if s.Key == language {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, "", errors.New(fmt.Sprintf("Language \"%s\" not found", language))
	}
	if s.Data.Subtitles[index].IsMachine {
		log.Println("Warning machine translation")
	}
	resp, err := utils.Request(s.Data.Subtitles[index].URL)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	fileType := filepath.Ext(resp.Request.URL.Path)
	switch fileType {
	case ".json":
		var subJson = new(BilibiliSubtitle)
		err := resp.Json(subJson)
		if err != nil {
			return nil, "", err
		}
		sub, err := jsonToSRT(subJson)
		if err != nil {
			return nil, "", err
		}
		return sub, ".srt", nil
	default:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
		return body, fileType, nil
	}
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

func jsonToSRT(subJson *BilibiliSubtitle) ([]byte, error) {
	var sub string
	var content string
	for i, s := range subJson.Body {
		if i != 0 || i == len(subJson.Body) {
			sub += "\n\n"
		}
		content = s.Content
		if s.Location != 2 {
			content = fmt.Sprintf("{\\an%d}%s", s.Location, content)
		}
		sub += fmt.Sprintf("%d\n%s --> %s\n%s", i+1, utils.SecondToTime(s.From), utils.SecondToTime(s.To), content)
	}
	return []byte(sub), nil
}
