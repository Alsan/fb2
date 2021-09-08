package cmd

import (
	"github.com/spf13/cobra"

	c "github.com/alsan/filebrowser/common"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/alsan/filebrowser/server/storage/bolt/importer"
)

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().String("old.database", "", "")
	upgradeCmd.Flags().String("old.config", "", "")
	_ = upgradeCmd.MarkFlagRequired("old.database")
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades an old configuration",
	Long: `Upgrades an old configuration. This command DOES NOT
import share links because they are incompatible with
this version.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		oldDB := mustGetString(flags, "old.database")
		oldConf := mustGetString(flags, "old.config")
		err := importer.Import(oldDB, oldConf, h.GetParam(flags, "database"))
		c.CheckErr(err)
	},
}
