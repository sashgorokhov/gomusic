package persistent_token

import (
	"github.com/sashgorokhov/gomusic/utils"
	"github.com/sashgorokhov/gomusic/utils/crypt"
	"github.com/sashgorokhov/govk"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type persistent_token map[string]govk.AuthInfo

func persistent_token_file_exists() bool {
	return utils.FileExists(utils.PERSISTENT_TOKEN_FILENAME)
}

func create_persistent_token_file() error {
	err := os.MkdirAll(path.Dir(utils.PERSISTENT_TOKEN_FILENAME), os.ModePerm)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(utils.PERSISTENT_TOKEN_FILENAME, []byte(""), os.ModePerm)
}

func read_persistent_token_file() (persistent_token, error) {
	var persistent_token_value persistent_token
	if persistent_token_file_exists() {
		contents, err := ioutil.ReadFile(utils.PERSISTENT_TOKEN_FILENAME)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(contents, &persistent_token_value)
		if err != nil {
			return nil, err
		}
	}
	return persistent_token_value, nil
}

func Add(login string, auth_info *govk.AuthInfo) error {
	var persistent_token_value persistent_token
	var err error

	if !persistent_token_file_exists() {
		err = create_persistent_token_file()
		if err != nil {
			return err
		}
		persistent_token_value = make(persistent_token)
	} else {
		persistent_token_value, err = read_persistent_token_file()
		if err != nil {
			return err
		}
	}
	encrypted, err := crypt.EncryptToken(auth_info.Access_token)
	if err != nil {
		return err
	}
	auth_info.Access_token = encrypted
	persistent_token_value[login] = *auth_info

	contents, err := yaml.Marshal(&persistent_token_value)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(utils.PERSISTENT_TOKEN_FILENAME, contents, os.ModePerm)
}

func Get(login string) (*govk.AuthInfo, bool) {
	if !persistent_token_file_exists() {
		return nil, false
	}
	persistent_token_value, err := read_persistent_token_file()
	if err != nil {
		return nil, false
	}
	v, ok := persistent_token_value[login]
	if !ok {
		return nil, false
	}
	if time.Now().After(v.Expires_at) {
		return nil, false
	}
	decrypted, err := crypt.DecryptToken(v.Access_token)
	if err != nil {
		return nil, false
	}
	v.Access_token = decrypted
	return &v, ok
}
