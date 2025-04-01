# RAG Service with OpenAI and Pinecone

This service implements Retrieval-Augmented Generation (RAG) using OpenAI for embeddings and completions, and Pinecone for vector storage and similarity search.

## Environment Setup

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Configure your environment variables in `.env`:
   ```
   OPENAI_API_KEY=your_api_key_here        # Required: OpenAI API key
   PINECONE_API_KEY=your_key_here          # Required: Pinecone API key
   PINECONE_INDEX_NAME=your_index_here     # Required: Pinecone index name
   OPENAI_MODEL=gpt-4                      # Optional: Default is gpt-4
   OPENAI_MAX_TOKENS=30                    # Optional: Default is 30
   ```

   Get your API keys from:
   - OpenAI API key: https://platform.openai.com/api-keys
   - Pinecone API key: https://app.pinecone.io/

### Pinecone Index Setup

Create a new index in Pinecone Console with these settings:
- Dimensions: 1536 (required for OpenAI ada-002 embeddings)
- Metric: Cosine
- Pod Type: starter (or higher based on your needs)

## Features

- Document ingestion with automatic embedding generation
- Semantic search using OpenAI embeddings
- RAG-based query answering using GPT-4
- Source attribution for answers
- Bulk JSON data upload utility
- RESTful API for OpenAI interactions

## Usage

### Running the API Server

1. Build the server:
   ```bash
   go build -o rag-server
   ```

2. Run the server:
   ```bash
   ./rag-server
   ```
   The server will start on port 8090 by default.

3. API Endpoint:

   **OpenAI Interaction**
   ```bash
   curl -X POST http://localhost:8090/openai \
     -H "Content-Type: application/json" \
     -d '{
       "query": "Your question here"
     }'
   ```

   Example response:
   ```json
   {
     "response": "AI-generated answer based on your query",
     "model": "gpt-4",
     "tokens_used": 150
   }
   ```

### JSON Uploader Tool

The project includes a command-line tool for uploading JSON chunks to Pinecone. This is useful for bulk data ingestion:

1. Build the uploader:
   ```bash
   go build -o json_uploader cmd/json_uploader/main.go
   ```

2. Prepare your JSON file:
   - Create a JSON file with an array of text chunks
   - Place it in the `data` directory
   - Example (`data/chunks.json`):
     ```json
     [
       "First text chunk to embed",
       "Second text chunk to embed",
       "Additional chunks..."
     ]
     ```

3. Run the uploader:
   ```bash
   ./json_uploader -input data/chunks.json
   ```

   The uploader will:
   - Read the JSON file
   - Generate embeddings using OpenAI's ada-002 model
   - Upload embeddings to your Pinecone index
   - Display progress as it processes chunks

### Development Examples

#### Using the RAG Service

```go
import "salihigde.com/go-openai-api-sample/services"

// Initialize the service
ragService, err := services.NewRAGService()
if err != nil {
    log.Fatal(err)
}

// Query the service
response, err := ragService.QueryWithRAG(context.Background(), &services.RAGRequest{
    Query: "Your question here",
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Answer: %s\n", response.Answer)
fmt.Printf("Sources: %v\n", response.Sources)
```

## Project Structure

```
.
├── cmd/
│   └── json_uploader/         # Bulk upload utility
├── config/                    # Configuration management
├── data/                     # Data files for upload
├── handlers/                 # HTTP request handlers
├── routers/                  # API route definitions
├── services/                 # Core business logic
│   ├── openai-service.go    # OpenAI service implementation
│   └── rag_service.go       # RAG service implementation
├── main.go                  # Main application entry point
├── .env                     # Environment variables (not in git)
├── .env.example            # Example environment variables
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
└── README.md              # Project documentation
```

## Security Notes

- Never commit your `.env` file to version control
- Keep your API keys secure and rotate them regularly
- Consider using environment variables in production instead of `.env` files