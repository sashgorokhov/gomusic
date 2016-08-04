package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sashgorokhov/gomusic/structs"
	"log"
	"fmt"
	"github.com/sashgorokhov/gomusic/formatters"
	"strconv"
	"os"
	"github.com/sashgorokhov/gomusic/utils"
)

var FriendsCommand = &cobra.Command{
	Use:   "friends",
	Short: "List friends",
	Long: `List friends`,
	Run: func(cmd *cobra.Command, args []string) {
		var friend_list structs.FriendResponse
		api, err := utils.Auth_by_flags(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		params := map[string]string{
			"count": strconv.Itoa(count),
			"offset": strconv.Itoa(offset),
			"fields": "domain",
			"order": "hints",
		}
		err = api.StructRequest("friends.get", params, &friend_list)
		if err != nil {
			log.Fatalln(err)
		}
		format, _ := cmd.Flags().GetString("format")
		if quiet {
			format = "id"
		}
		for _, v := range friend_list.Response.Items  {
			fmt.Println(formatters.Format_friend(&v, format))
		}
	},
}

func init() {
	FriendsCommand.Flags().IntVar(&offset, "offset", 0, "Offset")
	FriendsCommand.Flags().IntVarP(&count, "count", "c", 50, "How many friends to fetch. Specify -1 to show all available (offset also works here).")
	FriendsCommand.Flags().BoolVarP(&quiet, "quiet", "q", false, "Print only friends ids. Equal to --format=id")
	FriendsCommand.Flags().StringP("format", "f", formatters.Friend_format_default, "Print format. Available values: id, name. Mix it in desireble order.")
	utils.SetAuthFlags(FriendsCommand)
}
