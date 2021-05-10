package main

import "fmt"

func main() {
	// as like foreach
	nums := [3]int{1, 2, 3}
	for i, num := range nums {
		fmt.Println("index ", i, " = ", num)
	}
	fmt.Println("------------")
	// for loop
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------")
	// as like while loop
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
}
