package handler

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	n, err := w.Write([]byte("pong"))
	if err != nil {
		fmt.Printf("Could not sent a response: %v, %d", err, n)
	}
}
