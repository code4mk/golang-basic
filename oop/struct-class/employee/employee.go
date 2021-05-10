package employee

import (
	"fmt"
)

type Employee struct {
	FirstName string
	Age       int
	Phone     string
}

func New(FirstName string, Age int, Phone string) Employee {
	e := Employee{
		FirstName, Age, Phone,
	}
	return e
}

func (e Employee) Info() {
	fmt.Println(e)
}
