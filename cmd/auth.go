package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)


var authCmd = &cobra.Command{
	Use:   "auth -l login -p password",
	Short: "Authenticate user and print access token",
	Long: `Long help`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Api.Access_token)
	},
}

func init() {
	RootCmd.AddCommand(authCmd)
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
