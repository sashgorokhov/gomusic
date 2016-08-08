package utils

import (
	"log"
	"os/user"
	"path"
	"path/filepath"
)

const CLIENT_ID int = 5416726

var SCOPE = []string{"audio", "friends", "groups"}

var PERSISTENT_TOKEN_FILENAME, HOMEDIR, SETTINGSDIR, LOG_FILENAME string

func init() {
	user, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	HOMEDIR = filepath.ToSlash(user.HomeDir)
	SETTINGSDIR = path.Join(HOMEDIR, ".gomusic")
	PERSISTENT_TOKEN_FILENAME = path.Join(SETTINGSDIR, "persistent_tokens")
	LOG_FILENAME = path.Join(SETTINGSDIR, "gomusic.log")
}
