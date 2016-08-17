package cmd

import (
	"fmt"
	"github.com/sashgorokhov/gomusic/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"regexp"
)

var CleanlogCommand = &cobra.Command{
	Use:    "cleanlog",
	Short:  "Remove sensible data from logs",
	Long:   `Remove sensible data from logs and print them`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.FileExists(utils.LOG_FILENAME) {
			return
		}
		contents, err := ioutil.ReadFile(utils.LOG_FILENAME)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		regex, err := regexp.Compile("([Pp]assword[:=])(\\S+?)\\s|([Aa]ccess_token[:=])([\\w\\d]+)|([Aa]uth_secret[:=])([\\w\\s]+)\\s")
		fmt.Println(string(regex.ReplaceAll(contents, []byte("$1*hidden* "))))
	},
}
