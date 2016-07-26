package utils

import (
	"os/user"
	"log"
	"path"
	"path/filepath"
)

const CLIENT_ID int = 5416726
var SCOPE = []string{"audio"}


var PERSISTENT_TOKEN_FILENAME string


func init() {
	user, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	PERSISTENT_TOKEN_FILENAME = path.Join(filepath.ToSlash(user.HomeDir), ".gomusic", "persistent_tokens")
}
