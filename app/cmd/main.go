package main

import (
	"log"
	"shortify/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}
