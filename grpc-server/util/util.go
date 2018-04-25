package util

import (
	"log"
)

// Check is a helper func that checks deferred call on errors
func Check(fn func() error) {
	if err := fn(); err != nil {
		log.Printf("got error on deferred call: %v", err)
	}
}
