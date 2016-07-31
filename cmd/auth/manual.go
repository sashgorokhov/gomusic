package auth

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/sashgorokhov/govk"
	"github.com/sashgorokhov/gomusic/utils"
)


var ManualCommand = &cobra.Command{
	Use:   "manual",
	Short: "Manually get access token",
	Long: `Print login url to open it in browser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(govk.BuildLoginUrl(utils.CLIENT_ID, &utils.SCOPE))
	},
}

