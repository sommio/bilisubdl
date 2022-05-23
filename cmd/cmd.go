package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/K0ng2/bilisubdl/pkg/bilibili"
	"github.com/K0ng2/bilisubdl/utils"
	"github.com/spf13/cobra"
)

var (
	language string
	listSubs bool
	replace  bool
)

var RootCmd = &cobra.Command{
	Use:     "bilisubdl",
	Version: "1.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		if language == "" && !listSubs {
			log.Fatalln("No input language")
		}
		for _, s := range args {
			Run(s)
		}
	},
}

func init() {
	rootFlags := RootCmd.PersistentFlags()
	rootFlags.StringVarP(&language, "lang", "l", "", "Subtitle langague to download (e.g. en)")
	rootFlags.BoolVar(&listSubs, "list-subs", false, "List available subtitles language")
	rootFlags.BoolVarP(&replace, "replace", "r", false, "Replace existing subtitles")
}

func Run(id string) {
	var (
		title,
		subSRT,
		filename string
		episode *bilibili.Episode
		sub     *bilibili.Subtitle
	)
	info, err := bilibili.GetInfo(id)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Title:", info.Data.Season.Title)
	epList, err := bilibili.GetEpisodeList(id)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Total Episodes:", len(epList.Data.Sections[0].Episodes))

	title = utils.CleanText(info.Data.Season.Title)
	err = os.MkdirAll(title, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	for _, s := range epList.Data.Sections[0].Episodes {
		episode, err = bilibili.GetEpisode(s.EpisodeID)
		if err != nil {
			log.Println(err)
		}

		if listSubs {
			log.Println("Available subtitles for:", s.TitleDisplay)
			for _, s := range episode.Data.Subtitles {
				log.Println(s.LangKey, s.Lang)
			}
			continue
		}

		filename = filepath.Join(title, fmt.Sprintf("%s.%s.srt", utils.CleanText(s.TitleDisplay), language))
		if _, err := os.Stat(filename); err == nil && !replace {
			log.Println("Already exists:", filename)
			continue
		}

		sub, err = episode.GetSubtitleJSON(language)
		if err != nil {
			log.Fatalln(err)
		}

		subSRT = bilibili.SubToSRT(*sub)
		err := utils.CreateSubFile(filename, subSRT)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Writing subtitle to:", filename)
	}
	log.Println("Finished Downloading:", info.Data.Season.Title)
}
