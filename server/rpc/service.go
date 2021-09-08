package rpc

import (
	"context"

	fb "github.com/alsan/filebrowser/proto"
	h "github.com/alsan/filebrowser/server/helpers"
)

var (
	Data    h.PythonData
	Session string
)

type Server struct {
	fb.UnimplementedFileBrowserRpcServiceServer
}

func (s *Server) Login(ctx context.Context, in *fb.LoginRequest) (*fb.LoginReply, error) {
	// username := in.GetUsername()
	// password := in.GetPassword()

	return &fb.LoginReply{
		Status: fb.ReplyStatus_Ok,
		Data: &fb.LoginReply_Token{
			Token: "abcde",
		},
	}, nil
}
