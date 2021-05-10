package main

import "oop/employee"

func main() {
	e := employee.Employee{
		FirstName: "Sam",
		Age:       21,
		Phone:     "01751255157",
	}
	e.Info()

	// var p employee.Employee
	// p.Info()

	// constructor
	var q = employee.New("kamal", 21, "01851255158")
	q.Info()
}
