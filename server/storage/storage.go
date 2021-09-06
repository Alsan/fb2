package storage

import (
	"github.com/filebrowser/filebrowser/server/auth"
	"github.com/filebrowser/filebrowser/server/settings"
	"github.com/filebrowser/filebrowser/server/share"
	"github.com/filebrowser/filebrowser/server/users"
)

// Storage is a storage powered by a Backend which makes the necessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    users.Store
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
