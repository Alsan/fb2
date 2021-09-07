package storage

import (
	"github.com/alsan/filebrowser/server/auth"
	"github.com/alsan/filebrowser/server/settings"
	"github.com/alsan/filebrowser/server/share"
	"github.com/alsan/filebrowser/server/users"
)

// Storage is a storage powered by a Backend which makes the necessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    users.Store
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
