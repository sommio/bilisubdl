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
	language  string
	listSubs  bool
	overwrite bool
)

var RootCmd = &cobra.Command{
	Use:     "bilisubdl",
	Run: func(cmd *cobra.Command, args []string) {
		if language == "" && !listSubs {
			log.Fatalln("No input language")
		}
		for _, s := range args {
			Run(s)
		}
	},
	Example: "bilisubdl 37738 1042594 -l th\nbilisubdl 37738 --list-subs",
}

func init() {
	rootFlags := RootCmd.PersistentFlags()
	rootFlags.StringVarP(&language, "language", "l", "", "Subtitle language to download (e.g. en)")
	rootFlags.BoolVar(&listSubs, "list-subs", false, "List available subtitles language")
	rootFlags.BoolVarP(&overwrite, "overwrite", "w", false, "Force overwrite downloaded subtitles")
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

	if !listSubs {
		title = utils.CleanText(info.Data.Season.Title)
		err = os.MkdirAll(title, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}

	out:
	for _, j := range epList.Data.Sections {
		if !listSubs {
			log.Println("Episode List:", j.EpListTitle)
		}
		for _, s := range j.Episodes {
			filename = filepath.Join(title, fmt.Sprintf("%s.%s.srt", utils.CleanText(s.TitleDisplay), language))
			if _, err := os.Stat(filename); err == nil && !overwrite {
				log.Println("Already exists:", filename)
				continue
			}

			episode, err = bilibili.GetEpisode(s.EpisodeID)
			if err != nil {
				log.Println(err)
			}

			if listSubs {
				log.Println("Available subtitles language")
				for _, s := range episode.Data.Subtitles {
					log.Println(s.LangKey, s.Lang)
				}
				break out
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
	}
	if !listSubs {
		log.Println("Finished Downloading:", info.Data.Season.Title)
	}
}
