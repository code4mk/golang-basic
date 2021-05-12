package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	JobType string
	Salary  int
}

type Contractual struct {
	JobType string
	Salary  int
}

func (p Permanent) CalculateSalary() int {
	return p.Salary
}

func (c Contractual) CalculateSalary() int {
	return c.Salary + 1000
}

func TotalExpense(s SalaryCalculator) {
	fmt.Println(s.CalculateSalary())
}

func main() {
	fmt.Println("welcome")

	empPermanent := Permanent{
		JobType: "permanent",
		Salary:  2000,
	}

	TotalExpense(empPermanent)

	empContarctual := Contractual{
		JobType: "contract",
		Salary:  5000,
	}

	TotalExpense(empContarctual)
}
