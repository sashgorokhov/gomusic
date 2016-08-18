package cmd

import (
	"fmt"
	"github.com/sashgorokhov/gomusic/cmd/utils/auth"
	"github.com/sashgorokhov/gomusic/formatters"
	"github.com/sashgorokhov/gomusic/structs"
	"github.com/sashgorokhov/gomusic/utils"
	"github.com/sashgorokhov/govk"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var skip_error, skip_exists bool
var destination string
var post_id string

func make_audio_filename(audio *structs.Audio) string {
	return path.Join(filepath.ToSlash(destination), formatters.Format_audio_filename(audio))
}

var DownloadCommand = &cobra.Command{
	Use:   "download [audio_id [audio_id [...]]]",
	Short: "Music download",
	Long:  `Music download`,
	Run: func(cmd *cobra.Command, args []string) {
		var audio_list structs.AudioResponse
		api, err := auth.Authenticate(cmd)
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
		if len(args) > 0 {
			params["audio_ids"] = strings.Join(args, ",")
		}
		if post_id == "" {
			err = api.StructRequest("audio.get", params, &audio_list)
		} else {
			type post_response struct {
				Response []structs.Post `json:"response"`
			}
			var posts_list post_response
			err = api.StructRequest("wall.getById", map[string]string{
				"posts": post_id,
			}, &posts_list)
			if err == nil {
				audio_list.Response.Items = GetPostAudios(&posts_list.Response[0])
			}
		}
		if err != nil {
			fmt.Println(err)
			if error_struct, ok := err.(govk.ResponseError); ok {
				fmt.Printf("%+v\n", error_struct.ErrorStruct)
			}
			os.Exit(2)
		}

		os.MkdirAll(destination, os.ModeDir)

		for _, audio := range audio_list.Response.Items {
			filename := make_audio_filename(&audio)
			_, file := path.Split(filename)
			if _, err := os.Stat(filename); err == nil && skip_exists {
				fmt.Printf("%s: File exists - Skipping\n", file)
				continue
			}
			err := utils.Download_file(audio.CleanUrl(), filename)
			if err != nil {
				if !skip_error {
					panic(err)
				} else {
					fmt.Printf("Error while downloading %s: %s", file, err)
				}
			}
		}
	},
}

func init() {
	DownloadCommand.Flags().IntVar(&offset, "offset", 0, "Offset")
	DownloadCommand.Flags().IntVarP(&count, "count", "c", 50, "How many audios to fetch. Specify -1 to show all available (offset also works here).")
	DownloadCommand.Flags().IntVar(&owner_id, "owner_id", 0, "Owner id")
	DownloadCommand.Flags().StringVar(&post_id, "post_id", "", "Post id")
	DownloadCommand.Flags().IntVar(&album_id, "album_id", 0, "Album id")
	DownloadCommand.Flags().BoolVar(&skip_error, "skip_error", true, "Continue downloading if error occured")
	DownloadCommand.Flags().BoolVar(&skip_exists, "skip_exists", true, "Do not download audio if it is already downloaded")
	DownloadCommand.Flags().StringVarP(&destination, "destination", "d", "", "Where to save downloads")
	auth.SetAuthFlags(DownloadCommand)

}
