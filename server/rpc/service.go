package rpc

import (
	fb "github.com/alsan/filebrowser/proto"
	h "github.com/alsan/filebrowser/server/helpers"
	"github.com/alsan/filebrowser/server/settings"
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
