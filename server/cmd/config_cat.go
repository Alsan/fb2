package cmd

import (
	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/spf13/cobra"
)

func init() {
	configCmd.AddCommand(configCatCmd)
}

var configCatCmd = &cobra.Command{
	Use:   "cat",
	Short: "Prints the configuration",
	Long:  `Prints the configuration.`,
	Args:  cobra.NoArgs,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		set, err := d.Store.Settings.Get()
		c.CheckErr(err)
		ser, err := d.Store.Settings.GetServer()
		c.CheckErr(err)
		auther, err := d.Store.Auth.Get(set.AuthMethod)
		c.CheckErr(err)
		printSettings(ser, set, auther)
	}, h.PythonConfig{}),
}
