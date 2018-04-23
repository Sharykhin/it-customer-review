package http

import (
	"net/http"

	"fmt"
	"os"

	"github.com/Sharykhin/it-customer-review/api/http/handler"
	"github.com/Sharykhin/it-customer-review/api/middleware"
	"github.com/gorilla/mux"
)

// Handler is a main router for this service
func router() http.Handler {
	r := mux.NewRouter()
	r.Handle("/ping", middleware.Chain(
		middleware.Chain(http.HandlerFunc(handler.Pong),
			middleware.PanicRecover,
			middleware.CORS,
		))).Methods("GET")
	r.Handle("/reviews/{id}", middleware.Chain(
		middleware.Chain(http.HandlerFunc(handler.Get),
			middleware.PanicRecover,
			middleware.CORS,
		))).Methods("GET")
	r.Handle("/reviews", middleware.Chain(
		middleware.Chain(http.HandlerFunc(handler.Create),
			middleware.PanicRecover,
			middleware.CORS,
		))).Methods("POST")
	r.Handle("/reviews/{id}", middleware.Chain(
		middleware.Chain(http.HandlerFunc(handler.Update),
			middleware.PanicRecover,
			middleware.CORS,
		))).Methods("PUT")
	r.Handle("/reviews", middleware.Chain(
		middleware.Chain(http.HandlerFunc(handler.Index),
			middleware.PanicRecover,
			middleware.CORS,
		))).Methods("GET")
	return r
}

//ListenAndServe listens to all income http requests
func ListenAndServe() error {
	address := os.Getenv("HTTP_ADDRESS")
	fmt.Printf("Server is listening on %s\n", address)
	return http.ListenAndServe(address, router())
}
