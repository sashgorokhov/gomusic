package main

import (
	"github.com/sashgorokhov/gomusic/cmd"
	"github.com/sashgorokhov/gomusic/utils"
)

func main() {
	cleanup_func := utils.SetupLogging()
	if cleanup_func != nil {
		defer cleanup_func()
	}
	cmd.Execute()
}
