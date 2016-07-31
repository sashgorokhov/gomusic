package utils

import (
	"github.com/sashgorokhov/govk"
	"os"
	"bufio"
	"github.com/spf13/cobra"
	"errors"
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


func SetAuthFlags(c *cobra.Command) {
	c.Flags().StringP("login", "l", "", "Login")
	c.Flags().StringP("password", "p", "", "Password")
	c.Flags().StringP("access_token", "a", "", "Access Token")
	c.Flags().Bool("reuse_token", true, "Reuse access token stored in credentials file at ~/.gomusic/persistent_tokens")
}


func CheckAuthFlags(c *cobra.Command) error {
	login, password, access_token := GetAuthFlags(c)
	if access_token == "" && (login == "" || password == "") {
		return errors.New("--access_token or --login and --password must be specified")
	}
	return nil
}


func GetAuthFlags(c *cobra.Command) (login, password, access_token string) {
	login, _ = c.Flags().GetString("login")
	password, _ = c.Flags().GetString("password")
	access_token, _ = c.Flags().GetString("access_token")
	return
}


func Auth_by_flags(c *cobra.Command) (*govk.Api, error) {
	err := CheckAuthFlags(c)
	if err != nil {
		return nil, err
	}
	login, password, access_token := GetAuthFlags(c)
	reuse_token, err := c.Flags().GetBool("reuse_token")
	if err != nil {
		return nil, err
	}
	switch  {
		case access_token != "" :
			return Auth_by_login_and_password(login, password, reuse_token)
		case login != "" && password != "":
			return Auth_by_login_and_password(login, password, reuse_token)
		default:
			return nil, errors.New("All credentials are empty")

	}
}
