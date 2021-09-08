package rpc

import (
	"context"

	fb "github.com/alsan/filebrowser/proto"
	h "github.com/alsan/filebrowser/server/helpers"
	u "github.com/alsan/filebrowser/server/users"
)

var (
	storage *h.PythonData
)

type Server struct {
	fb.UnimplementedFileBrowserRpcServiceServer
}

func SetStorage(d *h.PythonData) {
	storage = d
}

func errInvalidUser() *fb.LoginReply {
	return &fb.LoginReply{
		Status: fb.ReplyStatus_Failed,
		Data: &fb.LoginReply_Message{
			Message: "Incorrect username or password.",
		},
	}
}

func (s *Server) Login(ctx context.Context, in *fb.LoginRequest) (*fb.LoginReply, error) {
	username := in.GetUsername()
	password := in.GetPassword()

	user, err := storage.Store.Users.Get("", username)
	if err != nil {
		if err.Error() == "the resource does not exist" {
			return errInvalidUser(), nil
		}

		return nil, err
	}

	if u.CheckPwd(password, user.Password) {
		return &fb.LoginReply{
			Status: fb.ReplyStatus_Ok,
			Data: &fb.LoginReply_Token{
				Token: "abcde",
			},
		}, nil
	}

	return errInvalidUser(), nil
}
