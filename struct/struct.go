package main 

import "fmt"

type user struct {
	name string
	age int
}

func main() {

 // struct with feild name
 e := user{
    name: "kamal",
    age: 21,
 }
 
 // struct without feild name
 r := user{"jamal",24}
 
 //  anynoumous struct 
 
 employee := struct{
    name string
 }{
   name: "maruf",
 }
 
 fmt.Println(e)
 fmt.Println(r)
 fmt.Println(employee)
 
 
}