package client

import (
	"context"
	"log"
	"time"

	pb "github.com/alsan/filebrowser/proto"
	"google.golang.org/grpc"
)

type FilelistClient struct {
	service pb.FileBrowserRpcServiceClient
	token   string
	path    string
	filter  string
}

func NewFilelistClient(conn *grpc.ClientConn, token, path, filter string) *FilelistClient {
	service := pb.NewFileBrowserRpcServiceClient(conn)
	return &FilelistClient{service, token, path, filter}
}

func (client *FilelistClient) GetFileList() []string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.service.FileList(ctx, &pb.FileListRequest{
		Token:  client.token,
		Path:   client.path,
		Filter: &client.filter,
	})
	if err != nil {
		log.Fatalf("Unable to get file list: %v", err)
	}

	if res.Status != pb.ReplyStatus_Ok {
		log.Fatalf("%s", res.GetMessage())
	}

	return res.GetList().Item
}
