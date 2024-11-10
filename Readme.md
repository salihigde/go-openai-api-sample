# Go OpenAI Sample Project

## Description
This is a sample project written in Go that demonstrates how to integrate with the OpenAI API. It provides a basic framework for making API calls to OpenAI and serves as a starting point for building more complex endpoints and features. This project is ideal for developers who want to learn how to structure a Go application, manage environment variables, and handle HTTP requests effectively. Use this as a foundation to expand into a fully-featured API with custom logic and additional integrations.

---

## Getting Started

Follow these steps to set up and run the application locally.

### Prerequisites

- **Go**: Make sure Go is installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
- **OpenAI API Key**: Obtain an API key from OpenAI and set it as an environment variable.

### Setup

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/your-username/your-repo-name.git
    cd your-repo-name
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

    If the project has any additional dependencies (e.g., Gorilla Mux), install them with:

    ```bash
    go mod tidy
    ```

### Running the Application

Run the application with:

```bash
go run main.go
```

### Testing the API
Once the application is running, you can make a test call to the API endpoint using curl:
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
This project follows a modular structure for improved maintainability and scalability:
```graphql
your-repo-name/
├── main.go                   # Application entry point
├── handlers/
│   └── openai.go             # Handler for OpenAI endpoint
├── services/
│   └── openai_service.go     # Business logic and API interaction with OpenAI
├── routers/
│   └── router.go             # Defines and initializes application routes
├── go.mod                    # Go module file
└── go.sum                    # Dependency lock file
```

- Handlers: Contains HTTP handler functions for different endpoints.
- Services: Encapsulates business logic and interactions with external APIs.
- Routers: Manages routing setup for the application.

### Contributing

Feel free to submit issues, fork the repository, and make pull requests. Any contributions are highly appreciated!

1. Fork the project.
2. Create your feature branch (```git checkout -b feature/new-feature```).
3. Commit your changes (```git commit -m 'Add new feature'```).
4. Push to the branch (```git push origin feature/new-feature```).
5. Open a Pull Request.

### Acknowledgments

- [OpenAI](https://openai.com) for providing the API.
- [Gorilla Mux](https://github.com/gorilla/mux) for HTTP routing in Go.