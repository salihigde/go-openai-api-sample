package routers

import (
	"github.com/gorilla/mux"
	"salihigde.com/go-openai-api-sample/handlers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/openai", handlers.OpenAIHandler).Methods("POST")
	return router
}
