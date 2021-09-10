package cmd

import (
	"log"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	// cobra.OnInitialize()
	// cobra.MousetrapHelpText = ""

	// rootCmd.SetVersionTemplate("File Browser version {{printf \"%s\" .Version}}\n")

	// flags := rootCmd.Flags()

	// flags.StringP("token", "t", "", "token to be used for authenticating the request")
	// flags.StringP("server", "s", "localhost:8080", "server address")
	// flags.StringP("username", "u", "alsan", "login username")
	// flags.StringP("password", "p", "KyHS4s77t1", "login password")

}

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

		reply := doLogin(server, username, password)
		if reply.GetStatus() == fb.ReplyStatus_Ok {
			token = reply.GetToken()
		} else {
			log.Fatalf("Login failed, message: %s", reply.GetMessage())
		}
	}

	return token
}
