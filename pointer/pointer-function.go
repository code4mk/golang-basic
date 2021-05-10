package main

import "fmt"

func change(val *string) {
	*val = "karim"
}

func main() {
	var a string = "kamal"

	var b = &a

	change(b)

	fmt.Println(a)
}
