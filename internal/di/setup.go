package di

import (
	"database/sql"
	"github.com/aberyotaro/go-api-sandbox/internal/app/translate"

	"github.com/aberyotaro/go-api-sandbox/internal/app/user"
	"github.com/aberyotaro/go-api-sandbox/internal/infrastructure"
	"github.com/aberyotaro/go-api-sandbox/internal/usecase"
)

type Handlers struct {
	db               *sql.DB
	UserHandler      *user.Handler
	TranslateHandler *translate.Handler
}

func NewHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		UserHandler: user.NewHandler(
			usecase.NewUser(
				db,
				infrastructure.NewUser(),
			),
		),
		TranslateHandler: translate.NewHandler(
			infrastructure.NewTranslate(),
		),
		// add new handler here
	}
}
