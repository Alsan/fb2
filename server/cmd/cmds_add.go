package cmd

import (
	"strings"

	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/spf13/cobra"
)

func init() {
	cmdsCmd.AddCommand(cmdsAddCmd)
}

var cmdsAddCmd = &cobra.Command{
	Use:   "add <event> <command>",
	Short: "Add a command to run on a specific event",
	Long:  `Add a command to run on a specific event.`,
	Args:  cobra.MinimumNArgs(2), //nolint:gomnd
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		s, err := d.Store.Settings.Get()
		c.CheckErr(err)
		command := strings.Join(args[1:], " ")
		s.Commands[args[0]] = append(s.Commands[args[0]], command)
		err = d.Store.Settings.Save(s)
		c.CheckErr(err)
		printEvents(s.Commands)
	}, h.PythonConfig{}),
}
