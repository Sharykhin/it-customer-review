package main

import (
	"fmt"

	"log"

	"github.com/Sharykhin/it-customer-review/tone-analyzer/analyzer"
)

func main() {
	a := analyzer.Analyzer
	s, err := a.Analyze("I have some ideas but let's discuss them together.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Score", s)
}
