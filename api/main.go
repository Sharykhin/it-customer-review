package main

import (
	"log"

	"github.com/Sharykhin/it-customer-review/api/http"
)

func main() {
	log.Fatal(http.ListenAndServe())
}
