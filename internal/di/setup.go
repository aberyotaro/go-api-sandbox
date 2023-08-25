package di

import (
	"github.com/aberyotaro/sample_api/internal/app/user"
	"github.com/aberyotaro/sample_api/internal/infrastructure"
	"github.com/aberyotaro/sample_api/internal/usecase"
)

type Handlers struct {
	UserHandler *user.Handler
}

func NewHandlers() *Handlers {
	return &Handlers{
		UserHandler: user.NewHandler(usecase.NewUser(infrastructure.NewUser())),
	}
}
