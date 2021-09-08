package cmd

import (
	c "github.com/alsan/filebrowser/common"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	configCmd.AddCommand(configSetCmd)
	addConfigFlags(configSetCmd.Flags())
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Updates the configuration",
	Long: `Updates the configuration. Set the flags for the options
you want to change. Other options will remain unchanged.`,
	Args: cobra.NoArgs,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		flags := cmd.Flags()
		set, err := d.Store.Settings.Get()
		c.CheckErr(err)

		ser, err := d.Store.Settings.GetServer()
		c.CheckErr(err)

		hasAuth := false
		flags.Visit(func(flag *pflag.Flag) {
			switch flag.Name {
			case "baseurl":
				ser.BaseURL = mustGetString(flags, flag.Name)
			case "root":
				ser.Root = mustGetString(flags, flag.Name)
			case "socket":
				ser.Socket = mustGetString(flags, flag.Name)
			case "cert":
				ser.TLSCert = mustGetString(flags, flag.Name)
			case "key":
				ser.TLSKey = mustGetString(flags, flag.Name)
			case "address":
				ser.Address = mustGetString(flags, flag.Name)
			case "port":
				ser.Port = mustGetString(flags, flag.Name)
			case "log":
				ser.Log = mustGetString(flags, flag.Name)
			case "signup":
				set.Signup = mustGetBool(flags, flag.Name)
			case "auth.method":
				hasAuth = true
			case "shell":
				set.Shell = convertCmdStrToCmdArray(mustGetString(flags, flag.Name))
			case "branding.name":
				set.Branding.Name = mustGetString(flags, flag.Name)
			case "branding.disableExternal":
				set.Branding.DisableExternal = mustGetBool(flags, flag.Name)
			case "branding.files":
				set.Branding.Files = mustGetString(flags, flag.Name)
			}
		})

		getUserDefaults(flags, &set.Defaults, false)

		// read the defaults
		auther, err := d.Store.Auth.Get(set.AuthMethod)
		c.CheckErr(err)

		// check if there are new flags for existing auth method
		set.AuthMethod, auther = getAuthentication(flags, hasAuth, set, auther)

		err = d.Store.Auth.Save(auther)
		c.CheckErr(err)
		err = d.Store.Settings.Save(set)
		c.CheckErr(err)
		err = d.Store.Settings.SaveServer(ser)
		c.CheckErr(err)
		printSettings(ser, set, auther)
	}, h.PythonConfig{}),
}
