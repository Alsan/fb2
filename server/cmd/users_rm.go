package cmd

import (
	"fmt"

	c "github.com/alsan/filebrowser/common"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/spf13/cobra"
)

func init() {
	usersCmd.AddCommand(usersRmCmd)
}

var usersRmCmd = &cobra.Command{
	Use:   "rm <id|username>",
	Short: "Delete a user by username or id",
	Long:  `Delete a user by username or id`,
	Args:  cobra.ExactArgs(1),
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		username, id := parseUsernameOrID(args[0])
		var err error

		if username != "" {
			err = d.Store.Users.Delete(username)
		} else {
			err = d.Store.Users.Delete(id)
		}

		c.CheckErr(err)
		fmt.Println("user deleted successfully")
	}, h.PythonConfig{}),
}
