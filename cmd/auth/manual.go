package auth

import (
	"fmt"
	"github.com/sashgorokhov/gomusic/utils"
	"github.com/sashgorokhov/govk"
	"github.com/spf13/cobra"
)

var ManualCommand = &cobra.Command{
	Use:   "manual",
	Short: "Manually get access token",
	Long:  `Print login url to open it in browser. How to obtain access token, refer to: https://new.vk.com/dev/implicit_flow_user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(govk.BuildLoginUrl(utils.CLIENT_ID, &utils.SCOPE))
	},
}
