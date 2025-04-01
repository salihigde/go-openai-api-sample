package handlers

import (
	"encoding/json"
	"net/http"

	"salihigde.com/go-openai-api-sample/services"
)

type OpenAIRequest struct {
	Prompt string `json:"prompt"`
}

type OpenAIResponse struct {
	Response string `json:"response"`
}

type OpenAIHandler struct {
	openaiService *services.OpenAIService
}

// NewOpenAIHandler creates a new instance of OpenAIHandler
func NewOpenAIHandler() (*OpenAIHandler, error) {
	openaiService, err := services.NewOpenAIService()
	if err != nil {
		return nil, err
	}

	return &OpenAIHandler{
		openaiService: openaiService,
	}, nil
}

// HandleOpenAI handles OpenAI API requests
func (h *OpenAIHandler) HandleOpenAI(w http.ResponseWriter, r *http.Request) {
	var requestBody OpenAIRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	responseText, err := h.openaiService.CallOpenAI(r.Context(), requestBody.Prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := OpenAIResponse{
		Response: responseText,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
