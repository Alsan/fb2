package cmd

import (
	"fmt"

	svc "github.com/alsan/filebrowser/client/client"
	c "github.com/alsan/filebrowser/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	cmd := fileListCmd
	flags := cmd.Flags()

	rootCmd.AddCommand(cmd)
	setCommonFlags(flags)
	flags.StringP("path", "P", "/", "path to the list of files")
	flags.StringP("filter", "f", "", "file extension to filter files")
}

var fileListCmd = &cobra.Command{
	Use:   "filelist",
	Short: "Get a list of files from server",
	Long:  `Get a list of files from server with optional path and filter`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer(cmd)
		conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
		c.ExitIfError("Unable to connect to server, %v", err)
		defer conn.Close()

		token := getLoginToken(cmd, server)
		flags := cmd.Flags()
		path, _ := flags.GetString("path")
		filter, _ := flags.GetString("filter")
		client := svc.NewFilelistClient(conn, token, path, filter)
		for _, f := range client.GetFileList() {
			fmt.Println(f)
		}
	},
}
