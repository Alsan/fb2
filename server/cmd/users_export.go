package cmd

import (
	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/spf13/cobra"
)

func init() {
	usersCmd.AddCommand(usersExportCmd)
}

var usersExportCmd = &cobra.Command{
	Use:   "export <path>",
	Short: "Export all users to a file.",
	Long: `Export all users to a json or yaml file. Please indicate the
path to the file where you want to write the users.`,
	Args: jsonYamlArg,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		list, err := d.Store.Users.Gets("")
		c.CheckErr(err)

		err = marshal(args[0], list)
		c.CheckErr(err)
	}, h.PythonConfig{}),
}
