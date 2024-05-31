package di

import (
	"database/sql"

	"github.com/aberyotaro/go-api-sandbox/internal/app/user"
	"github.com/aberyotaro/go-api-sandbox/internal/infrastructure"
	"github.com/aberyotaro/go-api-sandbox/internal/usecase"
)

type Handlers struct {
	db          *sql.DB
	UserHandler *user.Handler
}

func NewHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		UserHandler: user.NewHandler(
			usecase.NewUser(
				db,
				infrastructure.NewUser(),
			),
		),
		// add new handler here
	}
}
