package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/alsan/fb2/server/settings"
	"github.com/alsan/fb2/server/users"
)

func init() {
	rulesCmd.AddCommand(rulesRmCommand)
	rulesRmCommand.Flags().Uint("index", 0, "index of rule to remove")
	_ = rulesRmCommand.MarkFlagRequired("index")
}

var rulesRmCommand = &cobra.Command{
	Use:   "rm <index> [index_end]",
	Short: "Remove a global rule or user rule",
	Long: `Remove a global rule or user rule. The provided index
is the same that's printed when you run 'rules ls'. Note
that after each removal/addition, the index of the
commands change. So be careful when removing them after each
other.

You can also specify an optional parameter (index_end) so
you can remove all commands from 'index' to 'index_end',
including 'index_end'.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.RangeArgs(1, 2)(cmd, args); err != nil { //nolint:gomnd
			return err
		}

		for _, arg := range args {
			if _, err := strconv.Atoi(arg); err != nil {
				return err
			}
		}

		return nil
	},
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		i, err := strconv.Atoi(args[0])
		c.CheckErr(err)
		f := i
		if len(args) == 2 { //nolint:gomnd
			f, err = strconv.Atoi(args[1])
			c.CheckErr(err)
		}

		user := func(u *users.User) {
			u.Rules = append(u.Rules[:i], u.Rules[f+1:]...)
			err := d.Store.Users.Save(u)
			c.CheckErr(err)
		}

		global := func(s *settings.Settings) {
			s.Rules = append(s.Rules[:i], s.Rules[f+1:]...)
			err := d.Store.Settings.Save(s)
			c.CheckErr(err)
		}

		runRules(d.Store, cmd, user, global)
	}, h.PythonConfig{}),
}
