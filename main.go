package main

import (
	"fmt"
	"net/http"

	"salihigde.com/go-openai-api-sample/routers"
)

func main() {
	router := routers.InitRoutes()
	fmt.Println("Server started on port 8090")
	if err := http.ListenAndServe(":8090", router); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
