package cmd

import (
	"fmt"
	"github.com/sashgorokhov/gomusic/cmd/utils/auth"
	"github.com/sashgorokhov/gomusic/formatters"
	"github.com/sashgorokhov/gomusic/structs"
	"github.com/sashgorokhov/govk"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func GetPostAudios(post *structs.Post) []structs.Audio {
	var audio_list []structs.Audio
	for _, copy_post := range post.Copy_history {
		for _, audio := range GetPostAudios(&copy_post) {
			audio_list = append(audio_list, audio)
		}
	}
	for _, attach := range post.Attachments {
		if attach.Type == "audio" {
			audio_list = append(audio_list, attach.Audio)
		}
	}
	return audio_list
}

var PostsCommand = &cobra.Command{
	Use:   "posts",
	Short: "List posts",
	Long:  `List posts`,
	Run: func(cmd *cobra.Command, args []string) {
		var posts_list structs.PostResponse
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
		err = api.StructRequest("wall.get", params, &posts_list)
		if err != nil {
			fmt.Println(err)
			if error_struct, ok := err.(govk.ResponseError); ok {
				fmt.Printf("%+v\n", error_struct.ErrorStruct)
			}
			os.Exit(2)
		}
		for _, post := range posts_list.Response.Items {
			audio_list := GetPostAudios(&post)
			if len(audio_list) == 0 {
				continue
			}
			if !quiet {
				fmt.Printf("Post: %v_%v\n", int(post.Owner_id), post.Id)
				if post.Text != "" {
					fmt.Printf("Text: %v\n", post.Text)
				} else {
					if len(post.Copy_history) > 0 {
						fmt.Printf("Text: %v\n", post.Copy_history[0].Text)
					}
				}
				fmt.Println("Audios:")
			}
			for _, audio := range audio_list {
				fmt.Println(formatters.Format_audio(&audio, format))
			}
			if !quiet {
				fmt.Println()
			}
		}
	},
}

func init() {
	PostsCommand.Flags().BoolVarP(&quiet, "quiet", "q", false, "TODO")
	PostsCommand.Flags().IntVar(&owner_id, "owner_id", 0, "Owner id")
	PostsCommand.Flags().IntVar(&offset, "offset", 0, "Offset")
	PostsCommand.Flags().IntVarP(&count, "count", "c", 50, "How many groups to fetch. Specify -1 to show all available (offset also works here).")
	PostsCommand.Flags().StringVarP(&format, "format", "f", formatters.Audio_format_default, "Print format. Available values: id, url, title. Mix it in desirable order.")
	auth.SetAuthFlags(PostsCommand)
}
