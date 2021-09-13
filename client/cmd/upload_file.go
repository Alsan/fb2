package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type uploadFileClient struct {
	service  fb.FileBrowserRpcServiceClient
	token    string
	path     string
	filename string
}

func init() {
	cmd := uploadFileCmd
	flags := cmd.Flags()

	rootCmd.AddCommand(uploadFileCmd)
	setCommonFlags(flags)
	flags.StringP("path", "P", "/", "path to the list of files")
}

var uploadFileCmd = &cobra.Command{
	Use:   "uploadfile <filename>",
	Short: "Upload file to server",
	Long:  `Upload file to server by specifing server, path and filename.`,
	Args:  cobra.MinimumNArgs(1), //nolint:gomnd
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer(cmd)
		conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
		c.ExitIfError("Unable to connect to server, %v", err)
		defer conn.Close()

		token := getLoginToken(cmd, server)
		flags := cmd.Flags()
		path, _ := flags.GetString("path")
		filename := args[0]

		log.Printf("token: %s, path: %s filename: %s", token, path, filename)
		client := newUploadFileClient(conn, token, path, filename)
		ok := client.uploadFile()
		if !ok {
			log.Fatalf("error uploading file: %s\n", filename)
		}

		log.Println("file upload complete")
	},
}

func newUploadFileClient(conn *grpc.ClientConn, token, path, filename string) *uploadFileClient {
	service := fb.NewFileBrowserRpcServiceClient(conn)
	return &uploadFileClient{service, token, path, filename}
}

func (client *uploadFileClient) uploadFile() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	file, err := os.Open(client.filename)
	c.CheckErr(err)
	defer file.Close()

	req := &fb.UploadFileRequest{
		Token: client.token,
		Data: &fb.UploadFileRequest_Metadata{
			Metadata: &fb.FileInfo{
				Path:     client.path,
				Filename: client.filename,
				Size:     c.GetFileSize(file),
				Checksum: c.GetFileChecksum(file),
			},
		},
	}

	stream, err := client.service.UploadFile(ctx)
	c.CheckErr(err)

	err = stream.Send(req)
	c.CheckErr(err)

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		c.CheckErr(err)
		if err == io.EOF {
			break
		}

		req := &fb.UploadFileRequest{
			Token: client.token,
			Data: &fb.UploadFileRequest_Content{
				Content: buffer[:n],
			},
		}

		err = stream.Send(req)
		c.CheckErr(err)
	}

	res, err := stream.CloseAndRecv()
	c.CheckErr(err)

	if res.Status != fb.ReplyStatus_Ok {
		panic(fmt.Sprintf("File upload failed, %s", res.GetMessage()))
	}

	fmt.Println("File upload complete")
	return true
}
