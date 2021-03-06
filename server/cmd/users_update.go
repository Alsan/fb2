package cmd

import (
	"encoding/hex"

	"github.com/spf13/cobra"

	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/alsan/fb2/server/settings"
	"github.com/alsan/fb2/server/users"
)

func init() {
	usersCmd.AddCommand(usersUpdateCmd)

	usersUpdateCmd.Flags().StringP("password", "p", "", "new password")
	usersUpdateCmd.Flags().StringP("username", "u", "", "new username")
	addUserFlags(usersUpdateCmd.Flags())
}

var usersUpdateCmd = &cobra.Command{
	Use:   "update <id|username>",
	Short: "Updates an existing user",
	Long: `Updates an existing user. Set the flags for the
options you want to change.`,
	Args: cobra.ExactArgs(1),
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		username, id := parseUsernameOrID(args[0])
		flags := cmd.Flags()
		password := c.MustGetString(flags, "password")
		newUsername := c.MustGetString(flags, "username")

		var (
			err  error
			user *users.User
		)

		if id != 0 {
			user, err = d.Store.Users.Get("", id)
		} else {
			user, err = d.Store.Users.Get("", username)
		}

		c.CheckErr(err)

		defaults := settings.UserDefaults{
			Scope:       user.Scope,
			Locale:      user.Locale,
			ViewMode:    user.ViewMode,
			SingleClick: user.SingleClick,
			Perm:        user.Perm,
			Sorting:     user.Sorting,
			Commands:    user.Commands,
		}
		getUserDefaults(flags, &defaults, false)
		user.Scope = defaults.Scope
		user.Locale = defaults.Locale
		user.ViewMode = defaults.ViewMode
		user.SingleClick = defaults.SingleClick
		user.Perm = defaults.Perm
		user.Commands = defaults.Commands
		user.Sorting = defaults.Sorting
		user.LockPassword = c.MustGetBool(flags, "lockPassword")

		if newUsername != "" {
			user.Username = newUsername
		}

		if password != "" {
			user.Password = hex.EncodeToString(c.Md5Pass(password))
		}

		err = d.Store.Users.Update(user)
		c.CheckErr(err)
		printUsers([]*users.User{user})
	}, h.PythonConfig{}),
}
