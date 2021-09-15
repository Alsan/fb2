package cmd

import (
	h "github.com/alsan/fb2/server/helpers"
	"github.com/spf13/cobra"
)

func init() {
	rulesCmd.AddCommand(rulesLsCommand)
}

var rulesLsCommand = &cobra.Command{
	Use:   "ls",
	Short: "List global rules or user specific rules",
	Long:  `List global rules or user specific rules.`,
	Args:  cobra.NoArgs,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		runRules(d.Store, cmd, nil, nil)
	}, h.PythonConfig{}),
}
