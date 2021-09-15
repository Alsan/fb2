package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	c "github.com/alsan/fb2/common"
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
		pwd := hex.EncodeToString(c.Md5Pass(args[0]))

		fmt.Println(pwd)
	},
}
