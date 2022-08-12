package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, s := range args {
			err := Run(s)
			if err != nil {
				return err
			}
		}
		return nil
	},
	Example: "bilisubdl 37738 1042594 -l th\nbilisubdl 37738 --list-subs",
	SilenceErrors: true,
}

func init() {
	rootFlags := RootCmd.PersistentFlags()
	rootFlags.StringVarP(&language, "language", "l", "", "Subtitle language to download (e.g. en)")
	rootFlags.StringVarP(&output, "output", "o", "./", "Set output")
	rootFlags.BoolVarP(&listSubs, "list-subs", "L", false, "List available subtitle language")
	rootFlags.BoolVarP(&overwrite, "overwrite", "w", false, "Force overwrite downloaded subtitles")
}

func Run(id string) error {
	var (
		title, filename, fileType string
		episode                   *bilibili.BilibiliEpisode
		sub                       []byte
		exist                     bool
	)
	info, err := bilibili.Info(id)
	if err != nil {
		return err
	}

	epList, err := bilibili.Episodes(id)
	if err != nil {
		return err
	}

	title = utils.CleanText(info.Data.Season.Title)
	err = os.MkdirAll(title, os.ModePerm)
	if err != nil {
		return err
	}

	for _, j := range epList.Data.Sections {
		for _, s := range j.Episodes {
			filename = filepath.Join(output, title, fmt.Sprintf("%s.%s", utils.CleanText(s.TitleDisplay), language))
			for _, k := range []string{".srt", ".ass"} {
				if _, err := os.Stat(filename + k); err == nil && !overwrite {
					log.Println("#", filename+k)
					exist = true
					continue
				}
			}

			if exist {
				exist = false
				continue
			}

			episode, err = bilibili.Episode(s.EpisodeID.String())
			if err != nil {
				log.Println(err)
			}

			if listSubs {
				fmt.Printf("%-10s Title\n", "Key")
				fmt.Println(strings.Repeat("-", 20))
				for _, s := range episode.Data.Subtitles {
					fmt.Printf("%-10s %s\n", s.Key, s.Title)
				}
				return nil
			}

			sub, fileType, err = episode.Subtitle(language)
			if err != nil {
				return err
			}

			err := utils.WriteFile(filename+fileType, sub, s.PublishTime)
			if err != nil {
				return err
			}
			log.Println("*", filename+fileType)
		}
	}
	return nil
}
