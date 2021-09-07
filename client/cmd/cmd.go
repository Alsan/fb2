package cmd

import "github.com/alsan/filebrowser/utils"

func Execute() {
	err := rootCmd.Execute()
	utils.ExitIfError("Unable to execute root command, %v", err)
}