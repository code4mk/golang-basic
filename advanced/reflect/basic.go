package main

import (
	"fmt"
	"reflect"
)

// https://golang.org/pkg/reflect/

func main() {
	var name string = "kamal"
	r := reflect.TypeOf(name)
	fmt.Println(r)
}
