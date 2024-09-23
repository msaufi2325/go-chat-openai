# Go Chatbot

This project is a simple conversational chatbot implemented in Go. It uses the OpenAI API to generate responses based on user input. The chatbot can be customized with different personalities by passing an argument when running the program.

## Features

- Loads environment variables from a `.env` file using `godotenv`.
- Connects to the OpenAI API using the `go-openai` library.
- Customizable chatbot personality.
- Interactive command-line interface with colored text.

## Setup

1. Clone the repository:
  ```sh
  git clone https://github.com/msaufi2325/go-chatbot.git
  cd go-chatbot
  ```

2. Create a `.env` file in the root directory and add your OpenAI API key:
  ```env
  OPENAI_API_KEY=your_openai_api_key
  ```

3. Install the dependencies:
  ```sh
  go mod tidy
  ```

## Usage

Run the chatbot with the default personality:
```sh
go run main.go