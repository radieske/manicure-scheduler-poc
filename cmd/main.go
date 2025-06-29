package main

import (
	"log"

	server "github.com/radieske/manicure-scheduler-poc/internal/infra/web/http"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
