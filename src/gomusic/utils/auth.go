package utils

import (
	"github.com/sashgorokhov/govk"
	"os"
	"bufio"
)


func Auth_by_access_token(access_token string) (*govk.Api, error) {
	if _, err := os.Stat(access_token); err == nil {
		file, err := os.Open(access_token)
    		if err != nil {
    		    	return nil, err
    		} else {
			defer file.Close()
			reader := bufio.NewReader(file)
			new_access_token, err := reader.ReadString('\n')
			if err != nil {
				return nil, err
			}
			if new_access_token != "" {
				access_token = new_access_token
			}
		}
	}
	return govk.NewApi(access_token), nil
}


func Auth_by_login_and_password(login, password string, reuse_token bool) (*govk.Api, error) {
	var auth_info *govk.AuthInfo
	if reuse_token {
		auth_info, _ = Get(login)
	}
	if auth_info == nil {
		var err error
		auth_info, err = govk.Authenticate(login, password, CLIENT_ID, &SCOPE)
		if err != nil {
			return nil, err
		}
		if reuse_token {
			Add(login, auth_info)
		}

	}
	return Auth_by_access_token(auth_info.Access_token)
}
