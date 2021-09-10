package cmd

import (
	"log"

	c "github.com/alsan/filebrowser/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "fb-cli",
	Short: "filebrowser command line utility",
	Long:  "filebrowser command line utility",
}

func setCommonFlags(flags *pflag.FlagSet) {
	flags.StringP("token", "t", "", "token to be used for authenticating the request")
	flags.StringP("server", "s", "localhost:8080", "server address")
	flags.StringP("username", "u", "", "login username")
	flags.StringP("password", "p", "", "login password")
}

func getServer(cmd *cobra.Command) string {
	return c.MustGetString(cmd.Flags(), "server")
}

func getLoginToken(cmd *cobra.Command, server string) string {
	var token string

	flags := cmd.Flags()
	token = c.MustGetString(flags, "token")
	if token == "" {
		username := c.MustGetString(flags, "username")
		if username == "" {
			username = c.GetUserInput("username")
		}
		password := c.MustGetString(flags, "password")
		if password == "" {
			password = c.GetUserPasswordInput()
		}

		password = encryptPassword(password)
		msg, ok := doLogin(server, username, password)
		if !ok {
			log.Fatalf("Login failed, message: %s", msg)
		}

		token = msg
	}

	return token
}
