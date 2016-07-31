package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/sashgorokhov/gomusic/structs"
	"log"
	"github.com/sashgorokhov/gomusic/formatters"
	"strconv"
	"github.com/sashgorokhov/govk"
	"github.com/sashgorokhov/gomusic/utils"
)

var offset, count, owner_id, album_id int
var quiet bool
var format string
var Api *govk.Api

var MusicCommand = &cobra.Command{
	Use:   "music",
	Short: "List music",
	Long: `List music`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		Api, err = utils.Auth_by_flags(cmd)
		if err != nil {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var audio_list structs.AudioResponse
		params := map[string]string{
			"count": strconv.Itoa(count),
			"offset": strconv.Itoa(offset),
		}
		if owner_id != 0 {
			params["owner_id"] = strconv.Itoa(owner_id)
		}
		if album_id != 0 {
			params["album_id"] = strconv.Itoa(album_id)
		}
		err := Api.StructRequest("audio.get", params, &audio_list)
		if err != nil {
			log.Fatalln(err)
		}
		for _, v := range audio_list.Response.Items  {
			fmt.Println(formatters.Format_audio(&v, format, quiet, replace_chars))
		}
	},
}

func init() {
	MusicCommand.Flags().IntVar(&offset, "offset", 0, "Offset")
	MusicCommand.Flags().IntVarP(&count, "count", "c", 50, "How many audios to fetch. Specify -1 to show all available (offset also works here).")
	MusicCommand.Flags().IntVar(&owner_id, "owner_id", 0, "Owner id")
	MusicCommand.Flags().IntVar(&album_id, "album_id", 0, "Album id")
	MusicCommand.Flags().BoolVarP(&quiet, "quiet", "q", false, "Print only audio ids")
	MusicCommand.Flags().StringVarP(&format, "format", "f", formatters.Audio_format_default, "Print format. Available values: id, url, title. Mix it in desireble order.")
	utils.SetAuthFlags(MusicCommand)
}
