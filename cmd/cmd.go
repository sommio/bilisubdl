package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/K0ng2/bilisubdl/pkg/bilibili"
	"github.com/K0ng2/bilisubdl/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	language  string
	output    string
	listSubs  bool
	overwrite bool
	timeline  string
	search    string
)

var RootCmd = &cobra.Command{
	Use: "bilisubdl [id] [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case timeline != "-":
			err := RunTimeline()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		case search != "":
			err := RunSearch()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		default:
			for _, s := range args {
				err := Run(s)
				if err != nil {
					fmt.Fprintln(os.Stderr, "ID:", s, err)
				}
			}
		}
	},
	Example: "bilisubdl 37738 1042594 -l th\nbilisubdl 37738 --list-subs\nbilisubdl --timeline=sun",
}

func init() {
	rootFlags := RootCmd.PersistentFlags()
	rootFlags.StringVarP(&language, "language", "l", "", "Subtitle language to download (e.g. en)")
	rootFlags.StringVarP(&output, "output", "o", "./", "Set output")
	rootFlags.BoolVarP(&listSubs, "list-subs", "L", false, "List available subtitle language")
	rootFlags.BoolVarP(&overwrite, "overwrite", "w", false, "Force overwrite downloaded subtitles")
	rootFlags.StringVarP(&search, "search", "s", "", "Search anime")
	rootFlags.StringVarP(&timeline, "timeline", "T", "-", "Show timeline (sun|mon|tue|wed|thu|fri|sat)")
	rootFlags.Lookup("timeline").NoOptDefVal = "today"
}

func Run(id string) error {
	var (
		title, filename, fileType string
		episode                   *bilibili.Episode
		sub                       []byte
		exist                     bool
	)
	info, err := bilibili.GetInfo(id)
	if err != nil {
		return err
	}

	epList, err := bilibili.GetEpisodes(id)
	if err != nil {
		return err
	}

	title = utils.CleanText(info.Data.Season.Title)
	err = os.MkdirAll(filepath.Join(output, title), os.ModePerm)
	if err != nil {
		return err
	}

	for _, j := range epList.Data.Sections {
		for _, s := range j.Episodes {
			filename = filepath.Join(output, title, fmt.Sprintf("%s.%s", utils.CleanText(s.TitleDisplay), language))
			for _, k := range []string{".srt", ".ass"} {
				if _, err := os.Stat(filename + k); err == nil && !overwrite {
					fmt.Println("#", filename+k)
					exist = true
					continue
				}
			}

			if exist {
				exist = false
				continue
			}

			episode, err = bilibili.GetEpisode(s.EpisodeID.String())
			if err != nil {
				fmt.Println(err)
			}

			if listSubs {
				table := NewTable([]string{"Key", "Lang"})
				for _, s := range episode.Data.Subtitles {
					table.Append([]string{s.Key, s.Title})
				}
				fmt.Println("Title:", info.Data.Season.Title)
				table.Render()
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
			fmt.Println("*", filename+fileType)
		}
	}
	return nil
}

func RunTimeline() error {
	tl, err := bilibili.GetTimeline()
	if err != nil {
		return err
	}
	table := NewTable(nil)
	for _, s := range tl.Data.Items {
		if timeline == "today" {
			if s.IsToday {
				timeline = s.DayOfWeek
			}
		}
		if s.DayOfWeek == strings.ToUpper(timeline) {
			for _, j := range s.Cards {
				table.Append([]string{j.SeasonID, j.Title, j.PubTimeText})
			}
			table.SetHeader([]string{"ID", fmt.Sprintf("Title (%s %s)", s.DayOfWeek, s.FullDateText), "Status"})
			break
		}
	}
	table.Render()
	return nil
}

func RunSearch() error {
	ss, err := bilibili.GetSearch(search)
	if err != nil {
		return err
	}
	table := NewTable([]string{"ID", "Title", "Status"})
	for _, s := range ss.Data[1].Items {
		table.Append([]string{s.SeasonID.String(), s.Title, s.IndexShow})
	}
	table.Render()
	return nil
}

func NewTable(header []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetHeader(header)
	return table
}
