package cmd

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"reflect"

	"github.com/spf13/cobra"

	c "github.com/alsan/fb2/common"
	"github.com/alsan/fb2/server/auth"
	h "github.com/alsan/fb2/server/helpers"
	"github.com/alsan/fb2/server/settings"
)

func init() {
	configCmd.AddCommand(configImportCmd)
}

type settingsFile struct {
	Settings *settings.Settings `json:"settings"`
	Server   *settings.Server   `json:"server"`
	Auther   interface{}        `json:"auther"`
}

var configImportCmd = &cobra.Command{
	Use:   "import <path>",
	Short: "Import a configuration file",
	Long: `Import a configuration file. This will replace all the existing
configuration. Can be used with or without unexisting databases.

If used with a nonexisting database, a key will be generated
automatically. Otherwise the key will be kept the same as in the
database.

The path must be for a json or yaml file.`,
	Args: jsonYamlArg,
	Run: h.Python(func(cmd *cobra.Command, args []string, d h.PythonData) {
		var key []byte
		if d.HadDB {
			settings, err := d.Store.Settings.Get()
			c.CheckErr(err)
			key = settings.Key
		} else {
			key = generateKey()
		}

		file := settingsFile{}
		err := unmarshal(args[0], &file)
		c.CheckErr(err)

		file.Settings.Key = key
		err = d.Store.Settings.Save(file.Settings)
		c.CheckErr(err)

		err = d.Store.Settings.SaveServer(file.Server)
		c.CheckErr(err)

		var rawAuther interface{}
		if filepath.Ext(args[0]) != ".json" { //nolint:goconst
			rawAuther = cleanUpInterfaceMap(file.Auther.(map[interface{}]interface{}))
		} else {
			rawAuther = file.Auther
		}

		var auther auth.Auther
		switch file.Settings.AuthMethod {
		case auth.MethodJSONAuth:
			auther = getAuther(auth.JSONAuth{}, rawAuther).(*auth.JSONAuth)
		case auth.MethodNoAuth:
			auther = getAuther(auth.NoAuth{}, rawAuther).(*auth.NoAuth)
		case auth.MethodProxyAuth:
			auther = getAuther(auth.ProxyAuth{}, rawAuther).(*auth.ProxyAuth)
		default:
			c.CheckErr(errors.New("invalid auth method"))
		}

		err = d.Store.Auth.Save(auther)
		c.CheckErr(err)

		printSettings(file.Server, file.Settings, auther)
	}, h.PythonConfig{AllowNoDB: true}),
}

func getAuther(sample auth.Auther, data interface{}) interface{} {
	authType := reflect.TypeOf(sample)
	auther := reflect.New(authType).Interface()
	bytes, err := json.Marshal(data)
	c.CheckErr(err)
	err = json.Unmarshal(bytes, &auther)
	c.CheckErr(err)
	return auther
}
