package main

import (
	"log"

	"github.com/Sharykhin/it-customer-review/grpc-server/handler"
)

func main() {
	log.Fatal(handler.ListenAndServe())
}
