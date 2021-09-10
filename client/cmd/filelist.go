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

func init() {
	cmd := fileListCmd
	flags := cmd.Flags()

	rootCmd.AddCommand(cmd)
	flags.StringP("token", "t", "", "token to be used for authenticating the request")
	flags.StringP("server", "s", "localhost:8080", "server address")
	flags.StringP("username", "u", "", "login username")
	flags.StringP("password", "p", "", "login password")
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
		var token, server string

		flags := cmd.Flags()
		server = c.MustGetString(flags, "server")
		token = c.MustGetString(flags, "token")
		if token == "" {
			username := c.MustGetString(flags, "username")
			if username == "" {
				username = c.GetUserInput("username")
			}
			password := c.MustGetString(flags, "password")
			if password == "" {
				password = c.GetUserPasswordInput()
			}

			password = encryptPassword(password)

			reply := doLogin(server, username, password)
			if reply.GetStatus() == fb.ReplyStatus_Ok {
				token = reply.GetToken()
			} else {
				log.Fatalf("Login failed, message: %s", reply.GetMessage())
			}
		}

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
