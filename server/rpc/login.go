package rpc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"strings"
	"time"

	fb "github.com/alsan/filebrowser/proto"
	u "github.com/alsan/filebrowser/server/users"
	"google.golang.org/grpc/peer"
)

func errInvalidUser() *fb.LoginReply {
	return &fb.LoginReply{
		Status: fb.ReplyStatus_Failed,
		Data: &fb.LoginReply_Message{
			Message: "Incorrect username or password.",
		},
	}
}

func generateToken() string {
	tokenBuffer := make([]byte, 8) //nolint:gomnd
	if _, err := rand.Read(tokenBuffer); err != nil {
		log.Panic(err)
	}

	return hex.EncodeToString(tokenBuffer)
}

func getRemoteIP(ctx context.Context) string {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		log.Panic("Unable to parse peer from context")
	}

	ip := strings.Split(peer.Addr.String(), ":")[0]
	return strings.Join(strings.Split(ip, "."), "")
}

func (s *Server) Login(ctx context.Context, in *fb.LoginRequest) (*fb.LoginReply, error) {
	user, err := storage.Store.Users.Get("", in.GetUsername())
	if err != nil {
		// user not found, but we don't want the client knows the exact reason for safety
		if err.Error() == "the resource does not exist" {
			return errInvalidUser(), nil
		}

		// unknown error reason, should be handle by client side
		return nil, err
	}

	// verify user password
	if u.CheckPwd(in.GetPassword(), user.Password) {
		token := generateToken()
		ip := getRemoteIP(ctx)

		// save the session with token and timestamp
		session[token+"-"+ip] = time.Now().Unix()

		// login success
		return &fb.LoginReply{
			Status: fb.ReplyStatus_Ok,
			Data: &fb.LoginReply_Token{
				Token: token, // return client token only for safety
			},
		}, nil
	}

	// password not match
	return errInvalidUser(), nil
}
