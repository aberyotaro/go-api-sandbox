package infrastructure

import "github.com/aberyotaro/sample_api/internal/entity"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func NewUser() *User {
	return &User{}
}

func (i *User) GetByID(id int) (*entity.User, error) {
	// todo: db access
	dao := &User{
		ID:        id,
		FirstName: "Taro",
		LastName:  "Yamada",
	}

	return &entity.User{
		ID:        dao.ID,
		FirstName: dao.FirstName,
		LastName:  dao.LastName,
	}, nil
}
