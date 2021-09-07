package main

import (
	"context"
	"log"
	"time"

	fb "github.com/alsan/filebrowser/proto"
	"github.com/alsan/filebrowser/utils"

	"google.golang.org/grpc"
)

const ADDRESS = "127.0.0.1:8080"

func main() {
	// setup a connection to the server
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())
	utils.ExitIfError("Unable to connect to server, %v", err)
	defer conn.Close()

	client := fb.NewFileBrowserRpcServiceClient(conn)

	// create connection context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Login(ctx, &fb.LoginRequest{
		Username: "alsan",
		Password: "xxxx",
	})

	utils.ExitIfError("Unable to login, %v", err)

	if resp.GetStatus() == fb.ReplyStatus_Ok {
		log.Printf("Login successful, Token: %s", resp.GetToken())
	} else {
		log.Printf("Login failed, message: %s", resp.GetMessage())
	}
}
