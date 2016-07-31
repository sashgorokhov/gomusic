package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Audio list",
	Long: `Audio list`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
