package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alsan/filebrowser/utils"
)

func init() {
	rootCmd.AddCommand(hashCmd)
}

var hashCmd = &cobra.Command{
	Use:   "hash <password>",
	Short: "Hashes a password",
	Long:  `Hashes a password using bcrypt algorithm.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pwd := utils.Md5Pass(args[0])

		fmt.Println(pwd)
	},
}
