package usecase

import "github.com/aberyotaro/go-api-sandbox/internal/entity"

type Translator interface {
	Translate(text string, toLang entity.Lang) (*entity.ChatGPTResponse, error)
}
