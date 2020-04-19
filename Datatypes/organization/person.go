package organization

import "fmt"

type Indentifiable interface {
	ID() string
}

type Person struct {
	firstname string
	lastname  string
	twitterhandler string
}

func NewPerson(firstname, lastname string) Person {
	return Person{
		firstname: firstname,
		lastname:  lastname,
	}
}

func (p *Person) fullname() string {
	return fmt.Sprintf("%s %s", p.firstname, p.lastname)
}

func (p *Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterhandler = handler
	} else if !strings.HasPrefix(handler, prefix: "@") {
		return errors.New("twitter handle must start with an @ symbol")
	}

	p.twitterhandler = handler
	return nil
}

func (p Person) TwitterHandler() string {
	return p.twitterhandler
}