package utils

import (
	"os"
	"io/ioutil"
	"github.com/sashgorokhov/govk"
	"encoding/json"
	"path"
)

type persistent_token map[string]govk.AuthInfo

func persistent_token_file_exists() bool {
	_, err := os.Stat(PERSISTENT_TOKEN_FILENAME)
	return err == nil
}


func create_persistent_token_file() error {
	err := os.MkdirAll(path.Dir(PERSISTENT_TOKEN_FILENAME), os.ModePerm)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(PERSISTENT_TOKEN_FILENAME, []byte("{}"), os.ModePerm)
}


func read_persistent_token_file() (persistent_token, error) {
	var persistent_token_value persistent_token
	if persistent_token_file_exists() {
		contents, err := ioutil.ReadFile(PERSISTENT_TOKEN_FILENAME)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(contents, &persistent_token_value)
		if err != nil {
			return nil, err
		}
	}
	return persistent_token_value, nil
}


func Add(login string, auth_info *govk.AuthInfo) error {
	var persistent_token_value persistent_token
	var err error

	if ! persistent_token_file_exists() {
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
	persistent_token_value[login] = *auth_info

	contents, err := json.Marshal(&persistent_token_value)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(PERSISTENT_TOKEN_FILENAME, contents, os.ModePerm)
}


func Get(login string) (*govk.AuthInfo, bool) {
	if ! persistent_token_file_exists() {
		return nil, false
	}
	persistent_token_value, err := read_persistent_token_file()
	if err != nil {
		return nil, false
	}
	v, ok := persistent_token_value[login]
	return &v, ok
}