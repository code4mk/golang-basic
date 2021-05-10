package main

import "fmt"

func main() {
	var num [3]int
	num[0] = 1
	num[1] = 2
	num[2] = 3
	fmt.Println(num)

	a := [3]int{1, 2, 3}
	fmt.Println(a)

	b := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(b)

	// array slice
	c := b[1:3]
	fmt.Println(c)
	// modify slice data which also change main array
	c[0] = 25
	fmt.Println(b)
}
