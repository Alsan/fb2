package bolt

import (
	"github.com/asdine/storm"

	"github.com/filebrowser/filebrowser/server/auth"
	"github.com/filebrowser/filebrowser/server/settings"
	"github.com/filebrowser/filebrowser/server/share"
	"github.com/filebrowser/filebrowser/server/storage"
	"github.com/filebrowser/filebrowser/server/users"
)

// NewStorage creates a storage.Storage based on Bolt DB.
func NewStorage(db *storm.DB) (*storage.Storage, error) {
	userStore := users.NewStorage(usersBackend{db: db})
	shareStore := share.NewStorage(shareBackend{db: db})
	settingsStore := settings.NewStorage(settingsBackend{db: db})
	authStore := auth.NewStorage(authBackend{db: db}, userStore)

	err := save(db, "version", 2) //nolint:gomnd
	if err != nil {
		return nil, err
	}

	return &storage.Storage{
		Auth:     authStore,
		Users:    userStore,
		Share:    shareStore,
		Settings: settingsStore,
	}, nil
}
