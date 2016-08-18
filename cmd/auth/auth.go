package auth

import (
	"fmt"
	"github.com/sashgorokhov/gomusic/cmd/utils/auth"
	"github.com/spf13/cobra"
	"os"
)

var AuthCommand = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate user and print access token",
	Long:  `Long help`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flag("reuse_token").Value.Set("false")
		api, err := auth.Authenticate(cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(api.Access_token)
	},
}

func init() {
	AuthCommand.AddCommand(ManualCommand)
	//AuthCommand.AddCommand(SetCommand)
	auth.SetAuthFlags(AuthCommand)
}
