package functions

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) Greet() string{
	return "Hello " + p.FirstName
}

func (p *Person) HasBirthday() {
	p.Age++
}