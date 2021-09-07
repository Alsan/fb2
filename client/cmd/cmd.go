package cmd

import "github.com/alsan/filebrowser/common"

func Execute() {
	err := rootCmd.Execute()
	common.ExitIfError("Unable to execute root command, %v", err)
}
