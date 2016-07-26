package cmd

import (
	"github.com/spf13/cobra"
	"gomusic/structs"
	"log"
	"strconv"
	"gomusic/utils"
)

var skip_error, skip_exists bool
var destination string

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Music download",
	Long: `Music download`,
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
		for i, v := range audio_list.Response.Items  {
			go utils.DownloadAudio(v)
		}
	},
}

func init() {
	musicCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().IntVar(&offset, "offset", 0, "Offset")
	downloadCmd.Flags().IntVarP(&count, "count", "c", 50, "How many audios to fetch. Specify -1 to show all available (offset also works here).")
	downloadCmd.Flags().IntVar(&owner_id, "owner_id", 0, "Owner id")
	downloadCmd.Flags().IntVar(&album_id, "album_id", 0, "Album id")
	downloadCmd.Flags().BoolVar(&skip_error, "skip_error", true, "Continue downloading if error occured")
	downloadCmd.Flags().BoolVar(&skip_exists, "skip_exists", true, "Do not download audio if it is already downloaded")
	downloadCmd.Flags().StringVarP(&destination, "destination", "d", "", "Where to save downloads")

}
