package main

import (
	"fmt"
	"log"
	"net/http"

	"salihigde.com/go-openai-api-sample/config"
	"salihigde.com/go-openai-api-sample/routers"
)

func main() {
	// Initialize router
	router, err := routers.NewRouter()
	if err != nil {
		log.Fatalf("Failed to initialize router: %v", err)
	}

	// Start server
	fmt.Printf("Server is running on port%s\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerPort, router))
}
