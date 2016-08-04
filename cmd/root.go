package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/sashgorokhov/gomusic/cmd/auth"
)

var RootCmd = &cobra.Command{
	Use:   "gomusic",
	Short: "Download music from vkontakte",
	Long: `Download music from vkontakte`,
}


func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}


func init() {
	RootCmd.AddCommand(auth.AuthCommand)
	RootCmd.AddCommand(AlbumsCommand)
	RootCmd.AddCommand(MusicCommand)
	RootCmd.AddCommand(FriendsCommand)
}