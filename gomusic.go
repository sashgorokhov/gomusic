package main

import (
	"github.com/sashgorokhov/gomusic/cmd"
	"github.com/sashgorokhov/gomusic/utils"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] != "cleanlog" {
		cleanup_func := utils.SetupLogging()
		if cleanup_func != nil {
			defer cleanup_func()
		}
	}
	cmd.Execute()
}
