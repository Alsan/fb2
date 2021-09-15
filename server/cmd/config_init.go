package cmd

import (
	"fmt"

	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/spf13/cobra"

	"github.com/alsan/fb2/server/settings"
)

func init() {
	configCmd.AddCommand(configInitCmd)
	addConfigFlags(configInitCmd.Flags())
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new database",
	Long: `Initialize a new database to use with File Browser. All of
this options can be changed in the future with the command
'filebrowser config set'. The user related flags apply
to the defaults when creating new users and you don't
override the options.`,
	Args: cobra.NoArgs,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		defaults := settings.UserDefaults{}
		flags := cmd.Flags()
		getUserDefaults(flags, &defaults, true)
		authMethod, auther := getAuthentication(flags)

		s := &settings.Settings{
			Key:        generateKey(),
			Signup:     c.MustGetBool(flags, "signup"),
			Shell:      convertCmdStrToCmdArray(c.MustGetString(flags, "shell")),
			AuthMethod: authMethod,
			Defaults:   defaults,
			Branding: settings.Branding{
				Name:            c.MustGetString(flags, "branding.name"),
				DisableExternal: c.MustGetBool(flags, "branding.disableExternal"),
				Files:           c.MustGetString(flags, "branding.files"),
			},
		}

		ser := &settings.Server{
			Address: c.MustGetString(flags, "address"),
			Socket:  c.MustGetString(flags, "socket"),
			Root:    c.MustGetString(flags, "root"),
			BaseURL: c.MustGetString(flags, "baseurl"),
			TLSKey:  c.MustGetString(flags, "key"),
			TLSCert: c.MustGetString(flags, "cert"),
			Port:    c.MustGetString(flags, "port"),
			Log:     c.MustGetString(flags, "log"),
		}

		err := d.Store.Settings.Save(s)
		c.CheckErr(err)
		err = d.Store.Settings.SaveServer(ser)
		c.CheckErr(err)
		err = d.Store.Auth.Save(auther)
		c.CheckErr(err)

		fmt.Printf(`
Congratulations! You've set up your database to use with File Browser.
Now add your first user via 'filebrowser users add' and then you just
need to call the main command to boot up the server.
`)
		printSettings(ser, s, auther)
	}, h.PythonConfig{NoDB: true}),
}
