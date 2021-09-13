package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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
		// defer conn.Close()

		token := getLoginToken(cmd, server)
		flags := cmd.Flags()
		path, _ := flags.GetString("path")
		filename := args[0]

		if !strings.HasPrefix(filename, "//") {
			filename, err = filepath.Abs("./" + filename)
			c.ExitIfError("unable to find absolute path to file: %v", err)
		}

		if !c.IsFileExist(filename) {
			c.ExitIfError("unable to get current working directory, %v", err)
		}

		// log.Printf("token: %s, path: %s, filename: %s", token, path, filename)
		client := newUploadFileClient(conn, token, path, filename)
		ok := client.uploadFile()
		if !ok {
			conn.Close()
			log.Fatalf("error uploading file: %s\n", filename)
		}

		conn.Close()
		log.Println("file upload complete")
	},
}

func newUploadFileClient(conn *grpc.ClientConn, token, path, filename string) *uploadFileClient {
	service := fb.NewFileBrowserRpcServiceClient(conn)
	return &uploadFileClient{service, token, path, filename}
}

func (client *uploadFileClient) uploadFile() bool {
	file, err := os.Open(client.filename)
	c.ExitIfError("unable to open upload file, error: %v", err)
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := client.service.UploadFile(ctx)
	c.ExitIfError("unable to open upload stream, %v", err)

	err = stream.Send(&fb.UploadFileRequest{
		Data: &fb.UploadFileRequest_Metadata{
			Metadata: &fb.MetaData{
				Token:    client.token,
				Path:     client.path,
				Filename: client.filename,
				Size:     c.GetFileSize(file),
				Checksum: c.GetFileChecksum(file),
			},
		},
	})
	c.ExitIfError("unable to send metadata, %v", err)

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	// reset file pointer to ensure file read begining
	file.Seek(0, io.SeekStart)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		c.ExitIfError("unable to read chunk, %v", err)

		err = stream.Send(&fb.UploadFileRequest{
			Data: &fb.UploadFileRequest_Content{
				Content: buffer[:n],
			},
		})
		c.ExitIfError("unable to send chunk, %s", err)
	}

	res, err := stream.CloseAndRecv()
	c.ExitIfError("unable to receive response, %s", err)

	if res.Status != fb.ReplyStatus_Ok {
		panic(fmt.Sprintf("File upload failed, %s", res.GetMessage()))
	}

	fmt.Println("File upload complete")
	return true
}
