package cmd

import (
	"encoding/hex"

	"github.com/spf13/cobra"

	c "github.com/alsan/filebrowser/common"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/alsan/filebrowser/server/users"
)

func init() {
	usersCmd.AddCommand(usersAddCmd)
	addUserFlags(usersAddCmd.Flags())
}

var usersAddCmd = &cobra.Command{
	Use:   "add <username> <password>",
	Short: "Create a new user",
	Long:  `Create a new user and add it to the database.`,
	Args:  cobra.ExactArgs(2), //nolint:gomnd
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		s, err := d.Store.Settings.Get()
		c.CheckErr(err)
		getUserDefaults(cmd.Flags(), &s.Defaults, false)

		password := hex.EncodeToString(c.Md5Pass(args[1]))

		user := &users.User{
			Username:     args[0],
			Password:     password,
			LockPassword: c.MustGetBool(cmd.Flags(), "lockPassword"),
		}

		s.Defaults.Apply(user)

		servSettings, err := d.Store.Settings.GetServer()
		c.CheckErr(err)
		// since getUserDefaults() polluted s.Defaults.Scope
		// which makes the Scope not the one saved in the db
		// we need the right s.Defaults.Scope here
		s2, err := d.Store.Settings.Get()
		c.CheckErr(err)

		userHome, err := s2.MakeUserDir(user.Username, user.Scope, servSettings.Root)
		c.CheckErr(err)
		user.Scope = userHome

		err = d.Store.Users.Save(user)
		c.CheckErr(err)
		printUsers([]*users.User{user})
	}, h.PythonConfig{}),
}
