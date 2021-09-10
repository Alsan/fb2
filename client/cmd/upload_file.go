package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cmd := uploadFileCmd
	flags := cmd.Flags()

	rootCmd.AddCommand(uploadFileCmd)
	setCommonFlags(flags)
	flags.StringP("path", "P", "/", "path to the list of files")
}

var uploadFileCmd = &cobra.Command{
	Use:   "uploadfile <filename>",
	Short: "Upload file to server",
	Long:  `Upload file to server by specifing server, path and filename.`,
	Args:  cobra.MinimumNArgs(1), //nolint:gomnd
	Run: func(cmd *cobra.Command, args []string) {
		// server := getServer(cmd)
		// token := getLoginToken(cmd, server)
	},
}
