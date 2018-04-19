package handler

import (
	"net/http"

	"github.com/Sharykhin/it-customer-review/api/middleware"
	"github.com/gorilla/mux"
)

// Handler is a main router for this service
func Handler() http.Handler {
	r := mux.NewRouter()
	r.Handle("/ping", middleware.Chain(
		middleware.Chain(http.HandlerFunc(pong),
			middleware.PanicRecover,
			middleware.CORS,
		))).Methods("GET")

	r.Handle("/reviews", middleware.Chain(
		middleware.Chain(http.HandlerFunc(create),
			middleware.PanicRecover,
			middleware.CORS,
		))).Methods("POST")
	return r
}
