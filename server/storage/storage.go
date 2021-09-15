package storage

import (
	"github.com/alsan/fb2/server/auth"
	"github.com/alsan/fb2/server/settings"
	"github.com/alsan/fb2/server/share"
	"github.com/alsan/fb2/server/users"
)

// Storage is a storage powered by a Backend which makes the necessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    users.Store
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
