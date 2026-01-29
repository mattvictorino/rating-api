package main

import (
	"log"

	"github.com/mattvictorino/rating-api/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Cannot start application, an error has occurred %s", err.Error())
	}
}
