package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aberyotaro/go-api-sandbox/internal/entity"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (i *User) GetByID(db *sql.DB, id int) (*entity.User, error) {
	user := &entity.User{}

	q := `SELECT id, first_name, last_name, created_at, updated_at, deleted_at FROM users WHERE id = ? AND deleted_at IS NULL`

	if err := db.QueryRow(q, id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		log.Println(err)
	}

	fmt.Println("user: ", *user)

	return user, nil
}
