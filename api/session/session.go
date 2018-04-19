package session

import (
	"os"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte(os.Getenv("SESSION_KEY"))
	// Store is a session manager
	Store = sessions.NewCookieStore(key)
)
