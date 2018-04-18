package middleware

import (
	"net/http"
	"os"
)

var (
	corsOrigin string
)

func init() {
	corsOrigin = os.Getenv("CORS_ORIGIN")
}

//CORS adds all the necessary headers for cors
func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", corsOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
		h.ServeHTTP(w, r)
	})
}
