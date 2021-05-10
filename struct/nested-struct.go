package main 

import "fmt"

type address struct {
	city string
}

type user struct {
	name string
	age int
	area address
}

func main() {

 // struct with feild name
 e := user{
    name: "kamal",
    age: 21,
    area: address{
     city: "dhaka",
    },
 }
 
 fmt.Println(e.area.city)
 
}