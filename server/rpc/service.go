package rpc

import (
	"context"
	"log"

	fb "github.com/alsan/filebrowser/proto"
)

type Server struct {
	fb.UnimplementedFileBrowserRpcServiceServer
}

func (s *Server) Login(ctx context.Context, in *fb.LoginRequest) (*fb.LoginReply, error) {
	log.Printf("Login request received: %s", in.GetUsername())

	return &fb.LoginReply{
		Status: fb.ReplyStatus_Ok,
		Data: &fb.LoginReply_Token{
			Token: "abcde",
		},
	}, nil
}
