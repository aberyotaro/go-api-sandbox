package entity

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	endpoint = "https://api.openai.com/v1/chat/completions"
)

type ChatGPTModel string

const (
	Gpt4o ChatGPTModel = "gpt-4o"
)

type ChatGPTRequest struct {
	Model    ChatGPTModel            `json:"model"`
	Messages []ChatGPTRequestMessage `json:"messages"`
}

type ChatGPTRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTResponse struct {
	Choices []ChatGPTResponseChoice `json:"choices"`
}

type ChatGPTResponseChoice struct {
	Message ChatGPTResponseMessage `json:"message"`
}

type ChatGPTResponseMessage struct {
	Content string `json:"content"`
}

func getOpenAiApiKey() (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("APIキーが設定されていません")
	}
	return apiKey, nil
}

func NewOpenAIRequest(model ChatGPTModel, message []ChatGPTRequestMessage) ChatGPTRequest {
	return ChatGPTRequest{
		Model:    model,
		Messages: message,
	}
}

func NewOpenAIHttpRequest(jsonData []byte) (*http.Request, error) {
	apiKey, err := getOpenAiApiKey()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
