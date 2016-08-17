package auth

import (
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/pquerna/otp/totp"
	"github.com/sashgorokhov/gomusic/cmd/utils/auth/persistent_token"
	"github.com/sashgorokhov/gomusic/utils"
	"github.com/sashgorokhov/govk"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	"time"
)

type AuthFlags struct {
	Login        string
	Password     string
	Access_token string
	Auth_code    string
	Auth_secret  string
	Cfile        string
	Reuse_token  bool
}

var logger = utils.GomusicLogger.WithField("prefix", "gomusic.auth")

func SetAuthFlags(c *cobra.Command) {
	// Basic auth
	c.Flags().StringP("login", "l", "", "Login")
	c.Flags().StringP("password", "p", "", "Password")
	// Manual auth
	c.Flags().String("access_token", "", "Raw access token")
	// Two-factor auth
	c.Flags().String("auth_code", "", "Auth code for two-factor auth")
	c.Flags().String("auth_secret", "", "Secret for automatic two-factor auth code generation")

	// Get flags from file
	c.Flags().String("cfile", "", "YAML file containing credentials for authentication")

	c.Flags().Bool("reuse_token", true, "Reuse access token stored in credentials file at ~/.gomusic/persistent_tokens")
}

func GetAuthFlags(cmd *cobra.Command) (*AuthFlags, error) {
	auth_flags := AuthFlags{}
	auth_flags.Login, _ = cmd.Flags().GetString("login")
	auth_flags.Password, _ = cmd.Flags().GetString("password")
	auth_flags.Access_token, _ = cmd.Flags().GetString("access_token")
	auth_flags.Auth_code, _ = cmd.Flags().GetString("auth_code")
	auth_flags.Auth_secret, _ = cmd.Flags().GetString("auth_secret")
	auth_flags.Reuse_token, _ = cmd.Flags().GetBool("reuse_token")
	auth_flags.Cfile, _ = cmd.Flags().GetString("cfile")
	logger.WithField("auth_flags", fmt.Sprintf("%+v", auth_flags)).Debugln("Got auth flags")

	if auth_flags.Cfile != "" {
		if !utils.FileExists(auth_flags.Cfile) {
			return nil, errors.New(auth_flags.Cfile + " does not exist")
		}
		b, err := ioutil.ReadFile(auth_flags.Cfile)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"cfile": auth_flags.Cfile,
				"err":   err,
			}).Errorln("Cant read supplied cfile")
			return nil, err
		}
		err = yaml.Unmarshal(b, &auth_flags)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"cfile":    auth_flags.Cfile,
				"err":      err,
				"contents": string(b),
			}).Errorln("Cant unmarshall supplied cfile")
			return nil, err
		}
		logger.WithFields(logrus.Fields{
			"auth_flags": fmt.Sprintf("%+v", auth_flags),
			"cfile":      auth_flags.Cfile,
		}).Debugln("Got auth flags after unmarshalling cfile")
	}
	return &auth_flags, nil
}

func AuthByAccessToken(access_token string) (*govk.Api, error) {
	logger.Infoln("Authenticating by access token")
	return govk.NewApi(access_token), nil
}

func GenerateAuthCode(auth_secret string) (string, error) {
	return totp.GenerateCode(strings.Join(strings.Fields(auth_secret), ""), time.Now())
}

func AuthByLoginAndPassword(login, password, auth_code, auth_secret string, reuse_token bool) (*govk.Api, error) {
	var err error
	logger := logger.WithField("login", login)
	logger.Infoln("Authenticating by login and password")
	var auth_info *govk.AuthInfo
	if reuse_token {
		auth_info, _ = persistent_token.Get(login)
	}
	if auth_info == nil {
		if auth_secret != "" && auth_code == "" {
			logger.Infoln("Found auth_secret, and auth_code is empty. Generating...")
			auth_code, err = GenerateAuthCode(auth_secret)
			if err != nil {
				logger.WithField("err", err).Errorln("Error generating auth code")
				return nil, err
			}
		}
		auth_info, err = govk.Authenticate(login, password, utils.CLIENT_ID, &utils.SCOPE, auth_code)
		if err != nil {
			return nil, err
		}
		logger.WithField("auth_info", auth_info).Info("Successfully authenticated")

		err = persistent_token.Add(login, auth_info)
		if err != nil {
			logger.WithField("err", err).Warningln("Error adding persistent token")
		}
	}
	return AuthByAccessToken(auth_info.Access_token)
}

func AuthByFlags(auth_flags *AuthFlags) (*govk.Api, error) {
	switch {
	case auth_flags.Access_token != "":
		{
			return AuthByAccessToken(auth_flags.Access_token)
		}
	case auth_flags.Login != "" && auth_flags.Password != "":
		{
			return AuthByLoginAndPassword(auth_flags.Login, auth_flags.Password, auth_flags.Auth_code, auth_flags.Auth_secret, auth_flags.Reuse_token)
		}
	default:
		{
			logger.Errorln("Flags --access_token or --login and --password are required.")
			return nil, errors.New("Flags --access_token or --login and --password are required.")
		}

	}
}

func Authenticate(cmd *cobra.Command) (*govk.Api, error) {
	auth_flags, err := GetAuthFlags(cmd)
	if err != nil {
		return nil, err
	}
	return AuthByFlags(auth_flags)
}
