package cmd

import (
	"github.com/spf13/cobra"

	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/alsan/fb2/server/users"
)

func init() {
	usersCmd.AddCommand(usersFindCmd)
	usersCmd.AddCommand(usersLsCmd)
}

var usersFindCmd = &cobra.Command{
	Use:   "find <id|username>",
	Short: "Find a user by username or id",
	Long:  `Find a user by username or id. If no flag is set, all users will be printed.`,
	Args:  cobra.ExactArgs(1),
	Run:   findUsers,
}

var usersLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all users.",
	Args:  cobra.NoArgs,
	Run:   findUsers,
}

var findUsers = h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
	var (
		list []*users.User
		user *users.User
		err  error
	)

	if len(args) == 1 {
		username, id := parseUsernameOrID(args[0])
		if username != "" {
			user, err = d.Store.Users.Get("", username)
		} else {
			user, err = d.Store.Users.Get("", id)
		}

		list = []*users.User{user}
	} else {
		list, err = d.Store.Users.Gets("")
	}

	c.CheckErr(err)
	printUsers(list)
}, h.PythonConfig{})
