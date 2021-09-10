package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uploadFileCmd)
}

var uploadFileCmd = &cobra.Command{
	Use:   "uploadfile <token> <server> <path> <filename>",
	Short: "Upload file to server",
	Long:  `Upload file to server by specifing server, path and filename.`,
	Args:  cobra.MinimumNArgs(1), //nolint:gomnd
	Run: func(c *cobra.Command, args []string) {

	},
}
