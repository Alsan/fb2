package bolt

import (
	"github.com/asdine/storm"

	"github.com/filebrowser/filebrowser/server/auth"
	"github.com/filebrowser/filebrowser/server/errors"
	"github.com/filebrowser/filebrowser/server/settings"
)

type authBackend struct {
	db *storm.DB
}

func (s authBackend) Get(t settings.AuthMethod) (auth.Auther, error) {
	var auther auth.Auther

	switch t {
	case auth.MethodJSONAuth:
		auther = &auth.JSONAuth{}
	case auth.MethodProxyAuth:
		auther = &auth.ProxyAuth{}
	case auth.MethodNoAuth:
		auther = &auth.NoAuth{}
	default:
		return nil, errors.ErrInvalidAuthMethod
	}

	return auther, get(s.db, "auther", auther)
}

func (s authBackend) Save(a auth.Auther) error {
	return save(s.db, "auther", a)
}