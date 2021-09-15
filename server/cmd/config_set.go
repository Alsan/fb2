package cmd

import (
	c "github.com/alsan/fb2/common"
	h "github.com/alsan/fb2/server/helpers"
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
				ser.BaseURL = c.MustGetString(flags, flag.Name)
			case "root":
				ser.Root = c.MustGetString(flags, flag.Name)
			case "socket":
				ser.Socket = c.MustGetString(flags, flag.Name)
			case "cert":
				ser.TLSCert = c.MustGetString(flags, flag.Name)
			case "key":
				ser.TLSKey = c.MustGetString(flags, flag.Name)
			case "address":
				ser.Address = c.MustGetString(flags, flag.Name)
			case "port":
				ser.Port = c.MustGetString(flags, flag.Name)
			case "log":
				ser.Log = c.MustGetString(flags, flag.Name)
			case "signup":
				set.Signup = c.MustGetBool(flags, flag.Name)
			case "auth.method":
				hasAuth = true
			case "shell":
				set.Shell = convertCmdStrToCmdArray(c.MustGetString(flags, flag.Name))
			case "branding.name":
				set.Branding.Name = c.MustGetString(flags, flag.Name)
			case "branding.disableExternal":
				set.Branding.DisableExternal = c.MustGetBool(flags, flag.Name)
			case "branding.files":
				set.Branding.Files = c.MustGetString(flags, flag.Name)
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
