package main

import "fmt"

type User struct {
	Name string
}

type Employee struct {
	ID string
	User
}

func main() {

	u := User{
		Name: "kamal",
	}

	e := Employee{
		"56",
		u,
	}

	fmt.Println(e)

}
