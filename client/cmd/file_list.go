package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type filelistClient struct {
	service fb.FileBrowserRpcServiceClient
	token   string
	path    string
	filter  string
}

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
		client := newFilelistClient(conn, token, path, filter)
		for _, f := range client.GetFileList() {
			fmt.Println(f)
		}
	},
}

func newFilelistClient(conn *grpc.ClientConn, token, path, filter string) *filelistClient {
	service := fb.NewFileBrowserRpcServiceClient(conn)
	return &filelistClient{service, token, path, filter}
}

func (client *filelistClient) GetFileList() []string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.service.FileList(ctx, &fb.FileListRequest{
		Token:  client.token,
		Path:   client.path,
		Filter: &client.filter,
	})
	if err != nil {
		log.Fatalf("Unable to get file list: %v", err)
	}

	if res.Status != fb.ReplyStatus_Ok {
		log.Fatalf("%s", res.GetMessage())
	}

	return res.GetList().Item
}
