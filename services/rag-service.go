package services

import (
	"context"
	"fmt"
	"time"

	pinecone "github.com/pinecone-io/go-pinecone/v3/pinecone"
	"google.golang.org/protobuf/types/known/structpb"
	"salihigde.com/go-openai-api-sample/config"
)

type RAGService struct {
	openAIService  *OpenAIService
	pineconeClient *pinecone.Client
	indexName      string
}

type RAGRequest struct {
	Query string `json:"query"`
}

type RAGResponse struct {
	Answer  string   `json:"answer"`
	Sources []string `json:"sources,omitempty"`
}

// NewRAGService creates a new instance of RAGService
func NewRAGService() (*RAGService, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	openAIService, err := NewOpenAIService()
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAI service: %v", err)
	}

	// Create Pinecone client
	pineconeClient, err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: cfg.PineconeAPIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Pinecone client: %v", err)
	}

	return &RAGService{
		openAIService:  openAIService,
		pineconeClient: pineconeClient,
		indexName:      cfg.PineconeIndex,
	}, nil
}

// GenerateEmbedding generates embeddings for the given text using OpenAI
func (s *RAGService) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	return s.openAIService.GenerateEmbedding(ctx, text)
}

// QueryWithRAG performs RAG-based query processing
func (s *RAGService) QueryWithRAG(ctx context.Context, request *RAGRequest) (*RAGResponse, error) {
	// Generate embedding for the query
	queryEmbedding, err := s.GenerateEmbedding(ctx, request.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to generate query embedding: %v", err)
	}

	// Get index description to get the host
	idx, err := s.pineconeClient.DescribeIndex(ctx, s.indexName)
	if err != nil {
		return nil, fmt.Errorf("failed to describe index: %v", err)
	}

	// Create index connection
	index, err := s.pineconeClient.Index(pinecone.NewIndexConnParams{
		Host: idx.Host,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create index connection: %v", err)
	}

	// Query Pinecone for similar vectors
	queryResp, err := index.QueryByVectorValues(ctx, &pinecone.QueryByVectorValuesRequest{
		Vector:          queryEmbedding,
		TopK:            5,
		IncludeValues:   true,
		IncludeMetadata: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query Pinecone: %v", err)
	}

	// Extract context from matches
	var context string
	var sources []string
	for _, match := range queryResp.Matches {
		if match.Vector.Metadata != nil {
			if textField := match.Vector.Metadata.Fields["text"]; textField != nil {
				context += textField.GetStringValue() + "\n"
			}
			if sourceField := match.Vector.Metadata.Fields["source"]; sourceField != nil {
				sources = append(sources, sourceField.GetStringValue())
			}
		}
	}

	// Generate response using OpenAI with system prompt
	systemPrompt := "You are a helpful assistant. Use the provided context to answer the question like you are me. If you cannot find the answer in the context, say so."
	userPrompt := fmt.Sprintf("Context:\n%s\n\nQuestion: %s", context, request.Query)

	answer, err := s.openAIService.CallOpenAIWithSystemPrompt(ctx, userPrompt, systemPrompt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate response: %v", err)
	}

	return &RAGResponse{
		Answer:  answer,
		Sources: sources,
	}, nil
}

// UpsertDocument adds or updates a document in the vector store
func (s *RAGService) UpsertDocument(ctx context.Context, text, source string) error {
	// Generate embedding for the document
	embedding, err := s.GenerateEmbedding(ctx, text)
	if err != nil {
		return fmt.Errorf("failed to generate document embedding: %v", err)
	}

	// Get index description to get the host
	idx, err := s.pineconeClient.DescribeIndex(ctx, s.indexName)
	if err != nil {
		return fmt.Errorf("failed to describe index: %v", err)
	}

	// Create index connection
	index, err := s.pineconeClient.Index(pinecone.NewIndexConnParams{
		Host: idx.Host,
	})
	if err != nil {
		return fmt.Errorf("failed to create index connection: %v", err)
	}

	// Create metadata
	metadata, err := structpb.NewStruct(map[string]interface{}{
		"text":   text,
		"source": source,
	})
	if err != nil {
		return fmt.Errorf("failed to create metadata: %v", err)
	}

	// Create vector
	vectors := []*pinecone.Vector{
		{
			Id:       fmt.Sprintf("doc-%d", time.Now().UnixNano()),
			Values:   &embedding,
			Metadata: metadata,
		},
	}

	// Upsert to Pinecone
	count, err := index.UpsertVectors(ctx, vectors)
	if err != nil {
		return fmt.Errorf("failed to upsert to Pinecone: %v", err)
	}

	if count != 1 {
		return fmt.Errorf("expected to upsert 1 vector, but upserted %d", count)
	}

	return nil
}
