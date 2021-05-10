package main 

import "fmt"

type user struct {
	name string
	age int
}

func main() {

 // struct with feild name
 e := &user{
    name: "kamal",
    age: 21,
 }
 
 fmt.Println((*e).name)
 
}