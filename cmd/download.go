package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sashgorokhov/gomusic/structs"
	"log"
	"strconv"
	"path/filepath"
	"path"
	"github.com/sashgorokhov/gomusic/formatters"
	"os"
	"fmt"
	"github.com/sashgorokhov/gomusic/utils"
)

var skip_error, skip_exists bool
var destination string


func make_audio_filename(audio *structs.Audio) string {
	return path.Join(filepath.ToSlash(destination), formatters.Format_audio_filename(audio, replace_chars))
}


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
		for _, v := range audio_list.Response.Items  {
			filename := make_audio_filename(&v)
			if _, err := os.Stat(filename); err == nil && skip_exists {
				fmt.Println("Skipping")
				continue
			}
			_, file := path.Split(filename)
			pb := utils.ProgressBar{Title:file}
			pb.Init()
			err := utils.Download_file(v.CleanUrl(), filename, pb.Update)
			pb.Finish()
			if err != nil && ! skip_error {
				panic(err)
			}
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
