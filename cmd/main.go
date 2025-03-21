package main

import (
	"log"
	"net/http"

	"go-update-server/internal/handlers"
)

func main() {
	http.HandleFunc("/update", handlers.UpdateHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}