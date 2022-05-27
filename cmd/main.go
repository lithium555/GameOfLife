package main

import (
	"log"

	"GameOfLife/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("could not run application: %v", err)
	}
}
