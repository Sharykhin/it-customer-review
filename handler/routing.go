package handler

import (
	"net/http"

	"github.com/Sharykhin/it-customer-review/middleware"
	"github.com/gorilla/mux"
)

// Handler is a main router for this service
func Handler() http.Handler {
	r := mux.NewRouter()
	r.Handle("/ping", middleware.Chain(middleware.Chain(http.HandlerFunc(pong), middleware.PanicRecover))).Methods("GET")
	return r
}
