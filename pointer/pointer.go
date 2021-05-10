package main

import "fmt"

func main() {
	var a string = "kamal"

	var b = &a

	*b = "jamal"

	fmt.Println(a)
}
