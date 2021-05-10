package main

import "fmt"

func main() {
	// variable declare
	var age int = 21
	var name string = "kamal"
	fmt.Println("name is ", name, ", age = ", age)

	// multiple variable
	var height, width int = 200, 300
	fmt.Println("height =", height, " and width = ", width)

	// shorthand variable
	email := "hiremostafa@gmail.com"
	fmt.Println("email: ", email)

	// constant
	const a int = 21
	fmt.Println(" const a = ", a)
}
