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

var GroupsCommand = &cobra.Command{
	Use:   "groups",
	Short: "List groups",
	Long:  `List groups`,
	Run: func(cmd *cobra.Command, args []string) {
		var group_list structs.GroupResponse
		api, err := auth.Authenticate(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		params := map[string]string{
			"count":    strconv.Itoa(count),
			"offset":   strconv.Itoa(offset),
			"extended": "1",
		}
		err = api.StructRequest("groups.get", params, &group_list)
		if err != nil {
			fmt.Println(err)
			if error_struct, ok := err.(govk.ResponseError); ok {
				fmt.Printf("%+v\n", error_struct.ErrorStruct)
			}
			os.Exit(2)
		}
		format, _ := cmd.Flags().GetString("format")
		if quiet {
			format = "id"
		}
		for _, v := range group_list.Response.Items {
			fmt.Println(formatters.Format_group(&v, format))
		}
	},
}

func init() {
	GroupsCommand.Flags().IntVar(&offset, "offset", 0, "Offset")
	GroupsCommand.Flags().IntVarP(&count, "count", "c", 50, "How many groups to fetch. Specify -1 to show all available (offset also works here).")
	GroupsCommand.Flags().BoolVarP(&quiet, "quiet", "q", false, "Print only groups ids. Equal to --format=id")
	GroupsCommand.Flags().StringP("format", "f", formatters.Group_format_default, "Print format. Available values: id, name. Mix it in desireble order.")
	auth.SetAuthFlags(GroupsCommand)
}
