package cmd

import (
	"github.com/spf13/cobra"
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
