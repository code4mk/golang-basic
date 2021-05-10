package main

import (
	"fmt"
	"reflect"
)

// https://golang.org/pkg/reflect/
// https://betterprogramming.pub/understand-reflect-in-go-24a68fcf1011

func main() {
	var name string = "kamal"
	r := reflect.TypeOf(name)
	fmt.Println(r)
}
