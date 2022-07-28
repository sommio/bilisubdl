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
	output    string
	listSubs  bool
	overwrite bool
)

var RootCmd = &cobra.Command{
	Use: "bilisubdl [id] [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range args {
			Run(s)
		}
	},
	Example: "bilisubdl 37738 1042594 -l th\nbilisubdl 37738 --list-subs",
}

func init() {
	rootFlags := RootCmd.PersistentFlags()
	rootFlags.StringVarP(&language, "language", "l", "", "Subtitle language to download (e.g. en)")
	rootFlags.StringVarP(&output, "output", "o", "./", "Set output")
	rootFlags.BoolVarP(&listSubs, "list-subs", "L", false, "List available subtitles language")
	rootFlags.BoolVarP(&overwrite, "overwrite", "w", false, "Force overwrite downloaded subtitles")
}

func Run(id string) {
	var (
		title, filename, sub string
		episode              *bilibili.BilibiliEpisode
	)
	info, err := bilibili.Info(id)
	if err != nil {
		log.Fatalln(err)
	}

	epList, err := bilibili.Episodes(id)
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

	for _, j := range epList.Data.Sections {
		for _, s := range j.Episodes {
			// name := s.ShortTitleDisplay
			// if s.TitleDisplay != "" {
			// 	name = fmt.Sprintf("%s - %s", s.ShortTitleDisplay, utils.CleanText(s.TitleDisplay))
			// }
			filename = filepath.Join(output, title, fmt.Sprintf("%s.%s.srt", s.TitleDisplay, language))
			if _, err := os.Stat(filename); err == nil && !overwrite {
				log.Println("#", filename)
				continue
			}

			episode, err = bilibili.Episode(s.EpisodeID.String())
			if err != nil {
				log.Println(err)
			}

			if listSubs {
				log.Println("Available subtitles language")
				for _, s := range episode.Data.Subtitles {
					log.Println(s.Key, s.Title)
				}
				return
			}

			sub, err = episode.Subtitle(language)
			if err != nil {
				log.Fatalln(err)
			}

			err := utils.WriteFile(filename, sub)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("*", filename)
		}
	}
}
