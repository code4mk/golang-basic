package main

import "fmt"

func main(){
	fmt.Println("hello golang")
	user(216)
}

func user(id int){
	fmt.Println("function call ", id)
}