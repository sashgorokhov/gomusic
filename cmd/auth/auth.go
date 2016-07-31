package auth

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/sashgorokhov/gomusic/utils"
)


var AuthCommand = &cobra.Command{
	Use:   "auth -l <login> -p <password>",
	Short: "Authenticate user and print access token",
	Long: `Long help`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		_, err = cmd.Flags().GetString("login")
		if err != nil {
			return err
		}
		_, err = cmd.Flags().GetString("password")
		if err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		login, err := cmd.Flags().GetString("login")
		password, err := cmd.Flags().GetString("password")
		api, err := utils.Auth_by_login_and_password(login, password, false)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(api.Access_token)
	},
}

func init() {
	AuthCommand.AddCommand(ManualCommand)
	AuthCommand.AddCommand(SetCommand)
	AuthCommand.Flags().StringP("login", "l", "", "Login")
	AuthCommand.Flags().StringP("password", "p", "", "Password")
}