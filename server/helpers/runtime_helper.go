package helpers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/asdine/storm"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	v "github.com/spf13/viper"

	c "github.com/alsan/filebrowser/common"
	"github.com/alsan/filebrowser/server/storage"
	"github.com/alsan/filebrowser/server/storage/bolt"
)

func GetParam(flags *pflag.FlagSet, key string) string {
	val, _ := GetParamB(flags, key)
	return val
}

// getParamB returns a parameter as a string and a boolean to tell if it is different from the default
//
// NOTE: we could simply bind the flags to viper and use IsSet.
// Although there is a bug on Viper that always returns true on IsSet
// if a flag is binded. Our alternative way is to manually check
// the flag and then the value from env/config/gotten by viper.
// https://github.com/spf13/viper/pull/331
func GetParamB(flags *pflag.FlagSet, key string) (string, bool) {
	value, _ := flags.GetString(key)

	// If set on Flags, use it.
	if flags.Changed(key) {
		return value, true
	}

	// If set through viper (env, config), return it.
	if v.IsSet(key) {
		return v.GetString(key), true
	}

	// Otherwise use default value on flags.
	return value, false
}

type CobraFunc func(cmd *cobra.Command, args []string)
type PythonFunc func(cmd *cobra.Command, args []string, data PythonData)

type PythonConfig struct {
	NoDB      bool
	AllowNoDB bool
}

type PythonData struct {
	HadDB bool
	Store *storage.Storage
}

func DbExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return stat.Size() != 0, nil
	}

	if os.IsNotExist(err) {
		d := filepath.Dir(path)
		_, err = os.Stat(d)
		if os.IsNotExist(err) {
			if err := os.MkdirAll(d, 0700); err != nil { //nolint:govet,gomnd
				return false, err
			}
			return false, nil
		}
	}

	return false, err
}

func Python(fn PythonFunc, cfg PythonConfig) CobraFunc {
	return func(cmd *cobra.Command, args []string) {
		data := PythonData{HadDB: true}

		path := GetParam(cmd.Flags(), "database")
		exists, err := DbExists(path)

		if err != nil {
			panic(err)
		} else if exists && cfg.NoDB {
			log.Fatal(path + " already exists")
		} else if !exists && !cfg.NoDB && !cfg.AllowNoDB {
			log.Fatal(path + " does not exist. Please run 'filebrowser config init' first.")
		}

		data.HadDB = exists
		db, err := storm.Open(path)
		c.CheckErr(err)
		defer db.Close()
		data.Store, err = bolt.NewStorage(db)
		c.CheckErr(err)
		fn(cmd, args, data)
	}
}
