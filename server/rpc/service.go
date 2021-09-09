package rpc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/alsan/filebrowser/server/settings"
	u "github.com/alsan/filebrowser/server/users"
	"google.golang.org/grpc/peer"
)

var (
	storage    *h.PythonData
	serverConf *settings.Server
	session    = make(map[string]int64)
)

type Server struct {
	fb.UnimplementedFileBrowserRpcServiceServer
}

func SetStorage(d *h.PythonData) {
	storage = d
}

func SetServerConf(conf *settings.Server) {
	serverConf = conf
}

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

func getFilesWithoutFilter(path string) []string {
	var files []string

	err := filepath.Walk(serverConf.Root+path, func(curr string, info os.FileInfo, err error) error {
		files = append(files, curr)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func getFilesWithFilter(path, filter string) []string {
	var files []string
	var filterList = c.StrSlice(strings.Split(filter, ","))

	err := filepath.Walk(serverConf.Root+path, func(curr string, info os.FileInfo, err error) error {
		if filterList.Has(filepath.Ext(curr)) {
			files = append(files, curr)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return files

}
func getFileList(path, filter string) []string {
	if filter == "" {
		return getFilesWithoutFilter(path)
	}

	return getFilesWithFilter(path, filter)
}

func (s *Server) FileList(ctx context.Context, in *fb.FileListRequest) (*fb.FileListReply, error) {
	token := in.GetToken()
	ip := getRemoteIP(ctx)
	userToken := token + "-" + ip
	timestamp, ok := session[userToken]
	now := time.Now().Unix()

	if ok && now-timestamp < int64(time.Minute)*5 {
		// extend user sesion
		session[userToken] = now

		// return the file list
		return &fb.FileListReply{
			Status: fb.ReplyStatus_Ok,
			Data: &fb.FileListReply_List{
				List: &fb.List{
					Item: getFileList(in.GetPath(), in.GetFilter()),
				},
			},
		}, nil
	}

	// timeout or get the token from another ip
	return &fb.FileListReply{
		Status: fb.ReplyStatus_Failed,
		Data: &fb.FileListReply_Message{
			Message: "Invalid token, please login first",
		},
	}, nil
}
