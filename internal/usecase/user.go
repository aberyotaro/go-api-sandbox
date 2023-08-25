package usecase

import (
	"github.com/aberyotaro/sample_api/internal/entity"
	"github.com/aberyotaro/sample_api/internal/repository"
)

type User struct {
	repo repository.UserStore
}

func NewUser(repo repository.UserStore) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) GetUserByID(id int) (*entity.User, error) {
	return u.repo.GetByID(id)
}
