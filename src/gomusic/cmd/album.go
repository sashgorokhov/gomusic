package cmd

import (
	"github.com/spf13/cobra"
	"gomusic/structs"
	"log"
	"fmt"
	"gomusic/formatters"
	"strconv"
)

var albumsCmd = &cobra.Command{
	Use:   "albums",
	Short: "List albums",
	Long: `List albums`,
	Run: func(cmd *cobra.Command, args []string) {
		var album_list structs.AlbumResponse
		params := map[string]string{
			"count": strconv.Itoa(count),
			"offset": strconv.Itoa(offset),
		}
		if owner_id != 0 {
			params["owner_id"] = strconv.Itoa(owner_id)
		}
		err := Api.StructRequest("audio.getAlbums", params, &album_list)
		if err != nil {
			log.Fatalln(err)
		}
		for _, v := range album_list.Response.Items  {
			fmt.Println(formatters.Format_album(&v, format, quiet, replace_chars))
		}
	},
}

func init() {
	listCmd.AddCommand(albumsCmd)

	albumsCmd.Flags().IntVar(&offset, "offset", 0, "Offset")
	albumsCmd.Flags().IntVarP(&count, "count", "c", 50, "How many albums to fetch. Specify -1 to show all available (offset also works here).")
	albumsCmd.Flags().IntVar(&owner_id, "owner_id", 0, "Owner id")
	albumsCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Print only albums ids")
	albumsCmd.Flags().StringVarP(&format, "format", "f", formatters.Album_format_default, "Print format. Available values: id, title. Mix it in desireble order.")
}
