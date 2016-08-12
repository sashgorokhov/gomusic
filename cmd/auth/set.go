package auth

import (
	"fmt"
	"github.com/sashgorokhov/gomusic/cmd/utils/auth/persistent_token"
	"github.com/sashgorokhov/govk"
	"github.com/spf13/cobra"
	"os"
)

var SetCommand = &cobra.Command{
	Use:   "set [LOGIN] [ACCESS_TOKEN]",
	Short: "Set ACCESS_TOKEN for LOGIN",
	Long:  `Set ACCESS_TOKEN for LOGIN`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			auth_info := govk.AuthInfo{Access_token: args[1]}
			err := persistent_token.Add(args[0], &auth_info)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		} else {
			cmd.Usage()
			os.Exit(1)
		}

	},
}
