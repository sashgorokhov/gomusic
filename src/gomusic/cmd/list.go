package cmd

import (
	"github.com/spf13/cobra"
	"gomusic/structs"
	"log"
	"strconv"
)

var offset, count, owner_id int


var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Short help",
	Long: `Long help`,
	Run: func(cmd *cobra.Command, args []string) {
		var audio_list structs.AudioResponse
		params := map[string]string{
			"count": strconv.Itoa(count),
			"offset": strconv.Itoa(offset),
		}
		if owner_id != 0 {
			params["owner_id"] = strconv.Itoa(owner_id)
		}
		err := Api.StructRequest("audio.get", params, &audio_list)
		if err != nil {
			log.Fatalln(err)
		}
		for _, v := range audio_list.Response.Items  {
			log.Println(v)
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().IntVar(&offset, "offset", 0, "Offset")
	listCmd.Flags().IntVarP(&count, "count", "c", 50, "Count")
	listCmd.Flags().IntVar(&owner_id, "owner_id", 0, "Owner id")

}
