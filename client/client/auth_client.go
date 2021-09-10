package client

import (
	"context"
	"time"

	c "github.com/alsan/filebrowser/common"
	pb "github.com/alsan/filebrowser/proto"
	"google.golang.org/grpc"
)

// AuthClient is a client to call authentication RPC
type AuthClient struct {
	service  pb.FileBrowserRpcServiceClient
	username string
	password string
}

// NewAuthClient returns a new auth client
func NewAuthClient(conn *grpc.ClientConn, username, password string) *AuthClient {
	service := pb.NewFileBrowserRpcServiceClient(conn)
	return &AuthClient{service, username, password}
}

func GetConnection(server string) *grpc.ClientConn {
	conn, err := grpc.Dial(server, grpc.WithInsecure(), grpc.WithBlock())
	c.ExitIfError("Unable to connect to server, %v", err)
	defer conn.Close()

	return conn
}

// Login login user and returns the access token
func (client *AuthClient) Login() (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return false, "", err
	}

	if res.GetStatus() == pb.ReplyStatus_Ok {
		return true, res.GetToken(), nil
	}

	return false, res.GetMessage(), nil
}
