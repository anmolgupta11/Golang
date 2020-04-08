package models

type User struct {
	ID         int
	FirstName  string
	SecondName string
}

var (
	users  []*User
	nextID = 1
)
