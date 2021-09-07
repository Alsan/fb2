package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fb-cli",
	Short: "filebrowser command line utility",
	Long:  "filebrowser command line utility",
}
