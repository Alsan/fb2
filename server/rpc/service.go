package rpc

import (
	"context"
	"log"
	"time"

	fb "github.com/alsan/filebrowser/proto"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/alsan/filebrowser/server/settings"
)

const tokenTimeoutMessage string = "Token timeout or invalid token, please login"

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

func verifyToken(ctx context.Context, token string) bool {
	ip := getRemoteIP(ctx)
	userToken := token + "-" + ip
	timestamp, ok := session[userToken]
	now := time.Now().Unix()

	if ok && now-timestamp < int64(time.Minute)*5 {
		// extend user sesion
		session[userToken] = now

		return true
	}

	return false
}

func logError(err error) error {
	log.Print(err)
	return err
}
