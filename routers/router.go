package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"salihigde.com/go-openai-api-sample/handlers"
)

// NewRouter creates a new router instance
func NewRouter() (*mux.Router, error) {
	router := mux.NewRouter()

	// Create OpenAI handler
	openaiHandler, err := handlers.NewOpenAIHandler()
	if err != nil {
		return nil, err
	}

	ragHandler, err := handlers.NewRAGHandler()
	if err != nil {
		return nil, err
	}

	// Register routes
	router.HandleFunc("/openai", openaiHandler.HandleOpenAI).Methods(http.MethodPost)
	router.HandleFunc("/ragcv", ragHandler.QueryHandler).Methods(http.MethodPost)
	return router, nil
}
