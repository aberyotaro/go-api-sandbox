package repository

import (
	"database/sql"

	"github.com/aberyotaro/go-api-sandbox/internal/entity"
)

type UserStore interface {
	GetByID(db *sql.DB, id int) (*entity.User, error)
}
