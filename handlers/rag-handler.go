package handlers

import (
	"encoding/json"
	"net/http"

	"salihigde.com/go-openai-api-sample/services"
)

type RAGHandler struct {
	ragService *services.RAGService
}

func NewRAGHandler() (*RAGHandler, error) {
	ragService, err := services.NewRAGService()
	if err != nil {
		return nil, err
	}

	return &RAGHandler{
		ragService: ragService,
	}, nil
}

func (h *RAGHandler) QueryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request services.RAGRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.ragService.QueryWithRAG(r.Context(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *RAGHandler) UpsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type UpsertRequest struct {
		Text   string `json:"text"`
		Source string `json:"source"`
	}

	var request UpsertRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.ragService.UpsertDocument(r.Context(), request.Text, request.Source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}
