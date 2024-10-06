package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"salihigde.com/go-openai-api-sample/config"
)

type OpenAIRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}

const OpenAIRequestRole = "user"
const OpenAIRequestMaxTokens = 30

var httpClient = &http.Client{}

func CallOpenAI(prompt string) (string, error) {
	openAIRequest := OpenAIRequest{
		Model: config.OpenAIModel,
		Messages: []Message{
			{
				Role:    OpenAIRequestRole,
				Content: prompt,
			},
		},
		MaxTokens: OpenAIRequestMaxTokens,
	}

	reqBodyBytes, err := json.Marshal(openAIRequest)
	if err != nil {
		return "", err
	}

	apiKey := os.Getenv(config.OpenAIAPIKeyEnv)
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}

	req, err := http.NewRequest("POST", config.OpenAIEndpoint, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("OpenAI API returned status code %d: %s", resp.StatusCode, string(body))
	}

	var openaiResp OpenAIResponse
	err = json.Unmarshal(body, &openaiResp)
	if err != nil {
		return "", err
	}

	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("No response from OpenAI")
	}

	return openaiResp.Choices[0].Message.Content, nil
}
