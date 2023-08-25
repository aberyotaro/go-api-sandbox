package repository

import (
	"github.com/aberyotaro/sample_api/internal/entity"
)

type UserStore interface {
	GetByID(id int) (*entity.User, error)
}
