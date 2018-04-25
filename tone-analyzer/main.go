package main

import (
	"fmt"

	"log"

	"github.com/Sharykhin/it-customer-review/tone-analyzer/handler"
)

func main() {
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	log.Fatal(handler.ListenAndServe())
}
