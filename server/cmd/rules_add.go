package cmd

import (
	"regexp"

	"github.com/spf13/cobra"

	c "github.com/alsan/filebrowser/common"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/alsan/filebrowser/server/rules"
	"github.com/alsan/filebrowser/server/settings"
	"github.com/alsan/filebrowser/server/users"
)

func init() {
	rulesCmd.AddCommand(rulesAddCmd)
	rulesAddCmd.Flags().BoolP("allow", "a", false, "indicates this is an allow rule")
	rulesAddCmd.Flags().BoolP("regex", "r", false, "indicates this is a regex rule")
}

var rulesAddCmd = &cobra.Command{
	Use:   "add <path|expression>",
	Short: "Add a global rule or user rule",
	Long:  `Add a global rule or user rule.`,
	Args:  cobra.ExactArgs(1),
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		allow := c.MustGetBool(cmd.Flags(), "allow")
		regex := c.MustGetBool(cmd.Flags(), "regex")
		exp := args[0]

		if regex {
			regexp.MustCompile(exp)
		}

		rule := rules.Rule{
			Allow: allow,
			Regex: regex,
		}

		if regex {
			rule.Regexp = &rules.Regexp{Raw: exp}
		} else {
			rule.Path = exp
		}

		user := func(u *users.User) {
			u.Rules = append(u.Rules, rule)
			err := d.Store.Users.Save(u)
			c.CheckErr(err)
		}

		global := func(s *settings.Settings) {
			s.Rules = append(s.Rules, rule)
			err := d.Store.Settings.Save(s)
			c.CheckErr(err)
		}

		runRules(d.Store, cmd, user, global)
	}, h.PythonConfig{}),
}
