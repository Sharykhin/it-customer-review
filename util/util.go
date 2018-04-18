package util

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// GenerateRandomString returns securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}

// Check is a helper func that checks deferred call on errors
func Check(fn func() error) {
	if err := fn(); err != nil {
		log.Printf("got error on deferred call: %v", err)
	}
}
