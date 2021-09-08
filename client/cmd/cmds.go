package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdsCmd)
}

var cmdsCmd = &cobra.Command{
	Use:   "cmds",
	Short: "Command runner management utility",
	Long:  `Command runner management utility`,
	Args:  cobra.NoArgs,
}
