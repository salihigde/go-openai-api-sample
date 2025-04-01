package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var (
	once sync.Once
	cfg  *Config
)

// Config holds all configuration values
type Config struct {
	// OpenAI Configuration
	OpenAIAPIKey    string
	OpenAIModel     string
	OpenAIMaxTokens int

	// Pinecone Configuration
	PineconeAPIKey string
	PineconeIndex  string
}

// LoadConfig loads the configuration once and returns it
func LoadConfig() (*Config, error) {
	var loadErr error

	once.Do(func() {
		cfg = &Config{}

		// Load .env file from current directory
		if err := godotenv.Load(); err != nil {
			loadErr = fmt.Errorf("failed to load .env file: %v", err)
			return
		}

		// Load configuration values
		cfg.OpenAIAPIKey = os.Getenv("OPENAI_API_KEY")
		cfg.OpenAIModel = getEnvWithDefault("OPENAI_MODEL", openai.GPT4)
		cfg.OpenAIMaxTokens = getEnvAsIntWithDefault("OPENAI_MAX_TOKENS", 30)
		cfg.PineconeAPIKey = os.Getenv("PINECONE_API_KEY")
		cfg.PineconeIndex = os.Getenv("PINECONE_INDEX_NAME")

		// Validate required fields
		if err := validateConfig(cfg); err != nil {
			loadErr = err
			return
		}
	})

	if loadErr != nil {
		return nil, loadErr
	}

	return cfg, nil
}

// getEnvWithDefault returns environment variable value or default if not set
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsIntWithDefault returns environment variable as int or default if not set
func getEnvAsIntWithDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// validateConfig checks if all required fields are set
func validateConfig(cfg *Config) error {
	var missingVars []string

	if cfg.OpenAIAPIKey == "" {
		missingVars = append(missingVars, "OPENAI_API_KEY")
	}
	if cfg.PineconeAPIKey == "" {
		missingVars = append(missingVars, "PINECONE_API_KEY")
	}
	if cfg.PineconeIndex == "" {
		missingVars = append(missingVars, "PINECONE_INDEX_NAME")
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return nil
}

const ServerPort = ":8090"
