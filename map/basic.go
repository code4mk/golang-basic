package main

import "fmt"

func main() {
	user := make(map[string]string)
	user["name"] = "kamal"
	user["email"] = "hiremostafa@gmail.com"
	fmt.Println(user)

	employee := map[string]string{
		"id":     "em-12",
		"salary": "s12",
	}

	// map are reference type
	modified := employee
	modified["salary"] = "s13"

	for key, value := range employee {
		fmt.Println(key, " = ", value)
	}

	fmt.Println("id- ", employee["id"], "salary scale - ", employee["salary"])
}
