package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/sashgorokhov/gomusic/cmd/auth"
)

var replace_chars bool


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
	//RootCmd.PersistentFlags().BoolVar(&replace_chars, "replace_chars", true, "Only allow basic alphabet (rus+eng), digits and some signs.")
	//RootCmd.PersistentFlags().BoolVar(&reuse_token, "reuse_token", true, "Reuse access token")
}