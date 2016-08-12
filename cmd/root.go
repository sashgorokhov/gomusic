package cmd

import (
	"fmt"
	"github.com/sashgorokhov/gomusic/cmd/auth"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "gomusic",
	Short: "Download music from vkontakte",
	Long:  `Download music from vkontakte`,
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
	RootCmd.AddCommand(ListCommand)
	RootCmd.AddCommand(FriendsCommand)
	RootCmd.AddCommand(GroupsCommand)
	RootCmd.AddCommand(DownloadCommand)
}
