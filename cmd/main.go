package main

import (
	"log"
	"net/http"
)

func main() {
	// Composer
	userAdapter := Compose()

	// Create a router
	router := CreateRouter(userAdapter)

	// Start the server
	server := &http.Server{
		Addr:    ":6969",
		Handler: router,
	}

	log.Println("Server listening on port 6969")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
