package usecase

import (
	"database/sql"

	"github.com/aberyotaro/go-api-sandbox/internal/entity"
	"github.com/aberyotaro/go-api-sandbox/internal/repository"
)

type User struct {
	db   *sql.DB
	repo repository.UserStore
}

func NewUser(db *sql.DB, repo repository.UserStore) *User {
	return &User{
		db:   db,
		repo: repo,
	}
}

func (u *User) GetUserByID(id int) (*entity.User, error) {
	return u.repo.GetByID(u.db, id)
}
