package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/aberyotaro/go-api-sandbox/internal/entity"
	"io"
	"net/http"
)

type Translate struct{}

func NewTranslate() *Translate {
	return &Translate{}
}

func (t *Translate) Translate(text string, toLang entity.Lang) (*entity.ChatGPTResponse, error) {
	var (
		response *entity.ChatGPTResponse
		err      error
	)

	openAIReq, err := json.Marshal(entity.NewOpenAIRequest(entity.Gpt4o, []entity.ChatGPTRequestMessage{
		{Role: "system", Content: fmt.Sprintf("You are a helpful translator. translate the following text into %s .", toLang)},
		{Role: "user", Content: text},
	}))
	if err != nil {
		return nil, err
	}

	req, err := entity.NewOpenAIHttpRequest(openAIReq)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request error: %d - %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, fmt.Errorf("no response")
	}

	return response, err
}
