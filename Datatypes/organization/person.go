package organization

import "fmt"

type Indentifiable interface {
	ID() string
}

type Person struct {
	firstname string
	lastname  string
}

func NewPerson(firstname, lastname string) Person {
	return Person{
		firstname: firstname,
		lastname:  lastname,
	}
}

func (p Person) fullname() string {
	return fmt.Sprintf("%s %s", p.firstname, p.lastname)
}

func (p Person) ID() string {
	return "12345"
}
