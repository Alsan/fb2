package rpc

import (
	"context"
	"log"
	"time"

	fb "github.com/alsan/filebrowser/proto"
	"google.golang.org/grpc"
)

// AuthClient is a client to call authentication RPC
type AuthClient struct {
	service  fb.FileBrowserRpcServiceClient
	username string
	password string
}

// NewAuthClient returns a new auth client
func NewAuthClient(conn *grpc.ClientConn, username string, password string) *AuthClient {
	service := fb.NewFileBrowserRpcServiceClient(conn)
	return &AuthClient{service, username, password}
}

// Login login user and returns the access token
func (client *AuthClient) Login() (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &fb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		log.Fatalf("Unable to login: %v", err)
	}

	return res.GetToken(), true
}
