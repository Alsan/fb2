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
	rootCmd.AddCommand(fileListCmd)
}

func getOptionalArg(args []string, pos int, defaultValue string) string {
	if len(args) > pos {
		return args[pos]
	}

	return defaultValue
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
	Use:   "filelist <token> [server] [path] [filter]",
	Short: "Get a list of files from server",
	Long:  `Get a list of files from server with optional path and filter`,
	Args:  cobra.MinimumNArgs(1), //nolint:gomnd
	Run: func(c *cobra.Command, args []string) {
		token := args[0]
		server := getOptionalArg(args, 1, "localhost:8080")
		path := getOptionalArg(args, 2, "/")
		filter := getOptionalArg(args, 3, "")

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
