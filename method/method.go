package main 

import "fmt"

type user struct {
	name string
}

func main(){
	e := user{
	  name: "kamal",
	}
	e.employee()
}

func (e user) employee() {
	fmt.Println(e.name)
}