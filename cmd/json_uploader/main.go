package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"salihigde.com/go-openai-api-sample/services"
)

type JSONChunk struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// Parse command line flags
	inputFile := flag.String("input", "", "Path to the JSON file containing chunks")
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("Please provide an input file using -input flag")
	}

	// Read and parse JSON file
	chunks, err := readJSONFile(*inputFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Initialize RAG service
	ragService, err := services.NewRAGService()
	if err != nil {
		log.Fatalf("Failed to create RAG service: %v", err)
	}

	// Upload chunks to Pinecone
	ctx := context.Background()
	uploadChunksToPinecone(ctx, ragService, chunks)
}

func readJSONFile(filePath string) ([]JSONChunk, error) {
	// Find project root and construct full path
	projectRoot, err := findProjectRoot()
	if err != nil {
		return nil, fmt.Errorf("failed to find project root: %v", err)
	}

	fullPath := filepath.Join(projectRoot, filePath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var chunks []JSONChunk
	if err := json.Unmarshal(data, &chunks); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return chunks, nil
}

func uploadChunksToPinecone(ctx context.Context, ragService *services.RAGService, chunks []JSONChunk) {
	total := len(chunks)
	for i, chunk := range chunks {
		fmt.Printf("Uploading chunk %d/%d: %s\n", i+1, total, chunk.Title)
		if err := ragService.UpsertDocument(ctx, chunk.Content, chunk.Title); err != nil {
			log.Printf("Warning: Failed to upload chunk %d (%s): %v", i+1, chunk.Title, err)
			continue
		}
	}
	fmt.Println("Upload completed successfully!")
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root (no go.mod found)")
		}
		dir = parent
	}
}
