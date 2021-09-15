package cmd

import c "github.com/alsan/fb2/common"

func Execute() {
	err := rootCmd.Execute()
	c.ExitIfError("Unable to execute root command, %v", err)
}
