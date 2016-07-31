package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/sashgorokhov/govk"
	"github.com/sashgorokhov/gomusic/utils"
	"errors"
)

var replace_chars, reuse_token bool
var access_token string
var login, password string
var Api *govk.Api

var RootCmd = &cobra.Command{
	Use:   "gomusic",
	Short: "Download music from vkontakte",
	Long: ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		switch {
		case access_token != "":
			Api, err = utils.Auth_by_access_token(access_token)
			if err != nil {
				return err
			}
		case login != "" && password != "":
			Api, err = utils.Auth_by_login_and_password(login, password, reuse_token)
			if err != nil {
				return err
			}
		default:
			return errors.New("--access_token or --login and --pasword are required")
		}
		return nil
	},
//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}


func init() {
	RootCmd.PersistentFlags().BoolVar(&replace_chars, "replace_chars", true, "Only allow basic alphabet (rus+eng), digits and some signs.")
	RootCmd.PersistentFlags().BoolVar(&reuse_token, "reuse_token", true, "Reuse access token")
	RootCmd.PersistentFlags().StringVar(&access_token, "access_token", "", "Plain access token or path to file where first line is access token.")
	RootCmd.PersistentFlags().StringVarP(&login, "login", "l", "", "Login")
	RootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password")
}