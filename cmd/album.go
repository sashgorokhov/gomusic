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

var AlbumsCommand = &cobra.Command{
	Use:   "albums",
	Short: "List albums",
	Long:  `List albums`,
	Run: func(cmd *cobra.Command, args []string) {
		var album_list structs.AlbumResponse
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
		err = api.StructRequest("audio.getAlbums", params, &album_list)
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
		for _, v := range album_list.Response.Items {
			fmt.Println(formatters.Format_album(&v, format))
		}
	},
}

func init() {
	AlbumsCommand.Flags().IntVar(&offset, "offset", 0, "Offset")
	AlbumsCommand.Flags().IntVarP(&count, "count", "c", 50, "How many albums to fetch. Specify -1 to show all available (offset also works here).")
	AlbumsCommand.Flags().IntVar(&owner_id, "owner_id", 0, "Owner id")
	AlbumsCommand.Flags().BoolVarP(&quiet, "quiet", "q", false, "Print only albums ids. Equal to --format=id")
	AlbumsCommand.Flags().StringVarP(&format, "format", "f", formatters.Album_format_default, "Print format. Available values: id, title. Mix it in desireble order.")
	auth.SetAuthFlags(AlbumsCommand)
}
