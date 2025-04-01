package services

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"salihigde.com/go-openai-api-sample/config"
)

type OpenAIService struct {
	client    *openai.Client
	model     string
	maxTokens int
}

// NewOpenAIService creates a new instance of OpenAIService
func NewOpenAIService() (*OpenAIService, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	client := openai.NewClient(cfg.OpenAIAPIKey)
	return &OpenAIService{
		client:    client,
		model:     cfg.OpenAIModel,
		maxTokens: cfg.OpenAIMaxTokens,
	}, nil
}

// CallOpenAI sends a prompt to OpenAI and returns the response
func (s *OpenAIService) CallOpenAI(ctx context.Context, prompt string) (string, error) {
	return s.CallOpenAIWithSystemPrompt(ctx, prompt, "")
}

// CallOpenAIWithSystemPrompt sends a prompt to OpenAI with a custom system message
func (s *OpenAIService) CallOpenAIWithSystemPrompt(ctx context.Context, prompt string, systemPrompt string) (string, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	if systemPrompt != "" {
		messages = append([]openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
		}, messages...)
	}

	return s.CallOpenAIWithHistory(ctx, messages)
}

// CallOpenAIWithHistory sends a prompt to OpenAI with a conversation history
func (s *OpenAIService) CallOpenAIWithHistory(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	// Create the request
	req := openai.ChatCompletionRequest{
		Model:    s.model,
		Messages: messages,
	}

	// Only set MaxTokens for models that support it (not O3Mini)
	if s.model != "o3-mini" && s.model != "o3-mini-2025-01-31" {
		req.MaxTokens = s.maxTokens
	}

	resp, err := s.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned")
	}

	return resp.Choices[0].Message.Content, nil
}

// GenerateEmbedding generates embeddings for the given text using OpenAI
func (s *OpenAIService) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	resp, err := s.client.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Model: openai.AdaEmbeddingV2,
		Input: []string{text},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %v", err)
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}

	return resp.Data[0].Embedding, nil
}
