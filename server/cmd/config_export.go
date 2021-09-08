package cmd

import (
	c "github.com/alsan/filebrowser/common"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/spf13/cobra"
)

func init() {
	configCmd.AddCommand(configExportCmd)
}

var configExportCmd = &cobra.Command{
	Use:   "export <path>",
	Short: "Export the configuration to a file",
	Long: `Export the configuration to a file. The path must be for a
json or yaml file. This exported configuration can be changed,
and imported again with 'config import' command.`,
	Args: jsonYamlArg,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		settings, err := d.Store.Settings.Get()
		c.CheckErr(err)

		server, err := d.Store.Settings.GetServer()
		c.CheckErr(err)

		auther, err := d.Store.Auth.Get(settings.AuthMethod)
		c.CheckErr(err)

		data := &settingsFile{
			Settings: settings,
			Auther:   auther,
			Server:   server,
		}

		err = marshal(args[0], data)
		c.CheckErr(err)
	}, h.PythonConfig{}),
}
