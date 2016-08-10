package cmd

import (
	"fmt"
	"github.com/sashgorokhov/gomusic/formatters"
	"github.com/sashgorokhov/gomusic/structs"
	"github.com/sashgorokhov/gomusic/utils"
	"github.com/sashgorokhov/govk"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var offset, count, owner_id, album_id int
var quiet bool
var format string

var MusicCommand = &cobra.Command{
	Use:   "music",
	Short: "List music",
	Long:  `List music`,
	Run: func(cmd *cobra.Command, args []string) {
		var audio_list structs.AudioResponse
		api, err := utils.Auth_by_flags(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		params := map[string]string{
			"count":  strconv.Itoa(count),
			"offset": strconv.Itoa(offset),
		}
		if owner_id != 0 {
			params["owner_id"] = strconv.Itoa(owner_id)
		}
		if album_id != 0 {
			params["album_id"] = strconv.Itoa(album_id)
		}
		err = api.StructRequest("audio.get", params, &audio_list)
		if err != nil {
			fmt.Println(err)
			if error_struct, ok := err.(govk.ResponseError); ok {
				fmt.Printf("%+v\n", error_struct.ErrorStruct)
			}
			os.Exit(2)
		}
		if quiet {
			format = "id"
		}
		for _, v := range audio_list.Response.Items {
			fmt.Println(formatters.Format_audio(&v, format))
		}
	},
}

func init() {
	MusicCommand.AddCommand(DownloadCommand)
	MusicCommand.Flags().IntVar(&offset, "offset", 0, "Offset")
	MusicCommand.Flags().IntVarP(&count, "count", "c", 50, "How many audios to fetch. TODO: Specify -1 to show all available (offset also works here).")
	MusicCommand.Flags().IntVar(&owner_id, "owner_id", 0, "Owner id")
	MusicCommand.Flags().IntVar(&album_id, "album_id", 0, "Album id")
	MusicCommand.Flags().BoolVarP(&quiet, "quiet", "q", false, "Print only audio ids. Equal to --format=id")
	MusicCommand.Flags().StringVarP(&format, "format", "f", formatters.Audio_format_default, "Print format. Available values: id, url, title. Mix it in desirable order.")
	utils.SetAuthFlags(MusicCommand)
}
