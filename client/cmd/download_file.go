package cmd

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	c "github.com/alsan/fb2/common"
	fb "github.com/alsan/fb2/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type downloadFileClient struct {
	service    fb.FileBrowserRpcServiceClient
	token      string
	filename   string
	outputPath string
}

func init() {
	cmd := downloadFileCmd
	flags := cmd.Flags()

	rootCmd.AddCommand(cmd)
	setCommonFlags(flags)
	flags.StringP("output", "o", "./", "path to save the file")
}

var downloadFileCmd = &cobra.Command{
	Use:   "downloadfile <filename>",
	Short: "Download a file from server",
	Long:  "Download a file from server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer(cmd)
		conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
		c.ExitIfError("Unable to connect to server, %v", err)
		defer conn.Close()

		token := getLoginToken(cmd, server)
		filename := args[0]
		outputPath := c.MustGetString(cmd.Flags(), "output")

		client := newDownloadedClient(conn, token, filename, outputPath)
		if ok := client.downloadFile(); !ok {
			log.Fatalf("error download file: %s", filename)
		}

		log.Println("file downloaded successfully")
	},
}

func newDownloadedClient(conn *grpc.ClientConn, token, filename, outputPath string) *downloadFileClient {
	service := fb.NewFileBrowserRpcServiceClient(conn)
	return &downloadFileClient{service, token, filename, outputPath}
}

func (client *downloadFileClient) downloadFile() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := client.service.DownloadFile(ctx, &fb.DownloadFileRequest{
		Token:    client.token,
		Filename: client.filename,
	})
	if err != nil {
		log.Printf("Unable to send request, error: %v", err)
		return false
	}

	reply, err := stream.Recv()
	c.CheckErr(err)

	if reply.GetStatus() == fb.ReplyStatus_Failed {
		log.Printf("Download file failed: %v", reply.GetMessage())
		return false
	}

	info := reply.GetData().GetInfo()
	log.Printf("File info received: filename=%s, size=%d, checksum=%s", info.Filename, info.Size, info.Checksum)

	filename := filepath.Clean(filepath.Join(client.outputPath, filepath.Base(info.GetFilename())))
	log.Printf("File to be written to: %s", filename)

	file, err := os.Create(filename)
	c.CheckErr(err)
	defer file.Close()

	writter := bufio.NewWriter(file)

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		c.CheckErr(err)

		writter.Write(reply.GetData().GetContent())
	}

	writter.Flush()

	written := c.GetFileSize(file)
	if written != info.GetSize() {
		log.Printf("file size mismatch: %d", written)
		return false
	}

	checksum := c.GetFileChecksum(file)
	declared := info.GetChecksum()
	if checksum != declared {
		log.Printf("file checksum mismatch: %s", checksum)
		return false
	}

	return true
}
