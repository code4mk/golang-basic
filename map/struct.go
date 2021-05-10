package main

import "fmt"

type employee struct {
	name   string
	salary int
}

func main() {

	emp1 := employee{
		name:   "kamal",
		salary: 3000,
	}

	emp2 := employee{
		name:   "jamal",
		salary: 4000,
	}

	empInfo := map[string]employee{
		"emp1": emp1,
		"emp2": emp2,
	}

	fmt.Println("Length is = ", len(empInfo))

	for name, info := range empInfo {
		fmt.Println(name, " = ", info)
	}

}
