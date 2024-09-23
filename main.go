package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func bold(text string) string {
    return "\033[1m" + text + "\033[0m"
}

func blue(text string) string {
    return "\033[34m" + text + "\033[0m"
}

func red(text string) string {
    return "\033[31m" + text + "\033[0m"
}

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    apiKey := os.Getenv("OPENAI_API_KEY")
    client := openai.NewClient(apiKey)

    personality := "friendly and helpful"
    if len(os.Args) > 1 {
        personality = os.Args[1]
    }

    initialPrompt := fmt.Sprintf("You are a conversational chatbot. Your personality is: %s", personality)
    messages := []openai.ChatCompletionMessage{
        {Role: openai.ChatMessageRoleSystem, Content: initialPrompt},
    }

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(bold(blue("You: ")))
        userInput, _ := reader.ReadString('\n')
        userInput = strings.TrimSpace(userInput)
        messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: userInput})

        resp, err := client.CreateChatCompletion(
            context.Background(),
            openai.ChatCompletionRequest{
                Model:    "gpt-4o-mini",
                Messages: messages,
            },
        )

        if err != nil {
            fmt.Printf("ChatCompletion error: %v\n", err)
            break
        }

        messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: resp.Choices[0].Message.Content})
        fmt.Println(bold(red("Assistant: ")), resp.Choices[0].Message.Content)
    }
}
