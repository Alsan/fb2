package auth

import (
	"net/http"

	"github.com/alsan/fb2/server/users"
)

// Auther is the authentication interface.
type Auther interface {
	// Auth is called to authenticate a request.
	Auth(r *http.Request, s users.Store, root string) (*users.User, error)
	// LoginPage indicates if this auther needs a login page.
	LoginPage() bool
}
