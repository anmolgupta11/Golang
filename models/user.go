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

func GetUsers() []*Users {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}
