package handlers

import (
	"encoding/json"
	"net/http"

	"salihigde.com/go-openai-api-sample/services"
)

type OpenAIRequestBody struct {
	Prompt string `json:"prompt"`
}

type OpenAIResponseBody struct {
	Response string `json:"response"`
}

func OpenAIHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody OpenAIRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.Prompt == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	responseText, err := services.CallOpenAI(requestBody.Prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := OpenAIResponseBody{
		Response: responseText,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
