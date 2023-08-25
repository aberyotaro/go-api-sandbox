package entity

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
