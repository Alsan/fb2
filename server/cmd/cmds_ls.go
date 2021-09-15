package cmd

import (
	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/spf13/cobra"
)

func init() {
	cmdsCmd.AddCommand(cmdsLsCmd)
	cmdsLsCmd.Flags().StringP("event", "e", "", "event name, without 'before' or 'after'")
}

var cmdsLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all commands for each event",
	Long:  `List all commands for each event.`,
	Args:  cobra.NoArgs,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		s, err := d.Store.Settings.Get()
		c.CheckErr(err)
		evt := c.MustGetString(cmd.Flags(), "event")

		if evt == "" {
			printEvents(s.Commands)
		} else {
			show := map[string][]string{}
			show["before_"+evt] = s.Commands["before_"+evt]
			show["after_"+evt] = s.Commands["after_"+evt]
			printEvents(show)
		}
	}, h.PythonConfig{}),
}
