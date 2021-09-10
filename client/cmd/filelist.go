package cmd

import (
	"context"
	"fmt"
	"time"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
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

func doGetFileList(token, server, path, filter string) *fb.FileListReply {
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
	c.ExitIfError("Unable to connect to server, %v", err)
	defer conn.Close()
	client := fb.NewFileBrowserRpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if filter == "" {
		reply, err := client.FileList(ctx, &fb.FileListRequest{
			Token: token,
			Path:  path,
		})
		c.CheckErr(err)

		return reply
	}

	reply, err := client.FileList(ctx, &fb.FileListRequest{
		Token:  token,
		Path:   path,
		Filter: &filter,
	})
	c.CheckErr(err)

	return reply
}

var fileListCmd = &cobra.Command{
	Use:   "filelist",
	Short: "Get a list of files from server",
	Long:  `Get a list of files from server with optional path and filter`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer(cmd)
		token := getLoginToken(cmd, server)
		flags := cmd.Flags()
		path, _ := flags.GetString("path")
		filter, _ := flags.GetString("filter")
		reply := doGetFileList(token, server, path, filter)

		if reply.GetStatus() == fb.ReplyStatus_Ok {
			for _, file := range reply.GetList().GetItem() {
				fmt.Println(file)
			}
		} else {
			fmt.Printf("Error getting file list: %s\n", reply.GetMessage())
		}
	},
}
