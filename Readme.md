# Go OpenAI Sample Project

## Description
This is a sample project written in Go that demonstrates how to integrate with the OpenAI API. It provides a basic framework for making API calls to OpenAI and serves as a starting point for building more complex endpoints and features. This project is ideal for developers who want to learn how to structure a Go application, manage environment variables, and handle HTTP requests effectively. Use this as a foundation to expand into a fully-featured API with custom logic and additional integrations.

---

## Getting Started

Follow these steps to set up and run the application locally.

### Prerequisites

- **Go**: Make sure Go is installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
- **OpenAI API Key**: Obtain an API key from OpenAI by signing up at [OpenAI Platform](https://platform.openai.com/).

### Setup

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/salihigde/go-openai-api-sample.git
    cd go-openai-api-sample
    ```

2. **Set Up Environment Variables**:

    Set the `OPENAI_API_KEY` environment variable. This key is required to authenticate API requests to OpenAI.

    - **On macOS/Linux**:

      ```bash
      export OPENAI_API_KEY="your_openai_api_key_here"
      ```

    - **On Windows (Command Prompt)**:

      ```cmd
      set OPENAI_API_KEY=your_openai_api_key_here
      ```

    - **On Windows (PowerShell)**:

      ```powershell
      $env:OPENAI_API_KEY = "your_openai_api_key_here"
      ```

3. **Install Dependencies**:

    The project uses Go modules for dependency management. Install all required dependencies with:

    ```bash
    go mod tidy
    ```

### Running the Application

Run the application with:

```bash
go run main.go
```

By default, the server will start on port 8080.

### Testing the API
Once the application is running, you can test the API endpoint using curl:
```bash
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"prompt": "Hello, how are you?"}' \
     http://localhost:8080/openai
```

#### Expected Response
If everything is set up correctly, you should receive a response in JSON format that includes the AI's reply to your prompt:
```json
{
    "response": "I'm here to help! How can I assist you today?"
}
```

### Project Structure
This project follows a clean architecture pattern for improved maintainability and scalability:
```graphql
go-openai-api-sample/
├── main.go                   # Application entry point
├── config/                   # Configuration management
│   └── config.go            # Environment variables and app configuration
├── handlers/
│   └── openai.go            # HTTP handlers for OpenAI endpoints
├── services/
│   └── openai_service.go    # Business logic and OpenAI API integration
├── routers/
│   └── router.go            # HTTP routing setup
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
└── .gitignore              # Git ignore rules
```

#### Directory Structure Explanation
- **config**: Manages application configuration and environment variables
- **handlers**: Contains HTTP handler functions that process incoming requests
- **services**: Implements business logic and external API interactions
- **routers**: Defines API routes and middleware configuration

### Error Handling
The application includes robust error handling to manage common scenarios:
- Invalid API keys
- Network connectivity issues
- Rate limiting
- Malformed requests

### Best Practices
This project demonstrates several Go best practices:
- Environment-based configuration
- Modular project structure
- Clean separation of concerns
- Error handling patterns
- HTTP routing with Gorilla Mux

### Contributing

Contributions are welcome! Here's how you can help:

1. Fork the project
2. Create your feature branch (```git checkout -b feature/new-feature```)
3. Commit your changes (```git commit -m 'Add new feature'```)
4. Push to the branch (```git push origin feature/new-feature```)
5. Open a Pull Request

### License

This project is open source and available under the MIT License.

### Acknowledgments

- [OpenAI](https://openai.com) for providing the API
- [Gorilla Mux](https://github.com/gorilla/mux) for HTTP routing in Go

### Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.