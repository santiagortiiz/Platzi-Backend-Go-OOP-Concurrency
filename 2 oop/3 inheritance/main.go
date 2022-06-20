/*
Go does not implement inheritance, instead
it implements composition over inheritance.
*/
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) GetName() {
	fmt.Printf("%v is %v years old", p.name, p.age)
}

type Employee struct {
	id       int
	vacation bool
}

type FullTimeEmployee struct {
	// Anonymous fields define the composition
	Person
	Employee
}

func main() {
	fte := FullTimeEmployee{}
	fte.id = 1
	fte.name = "santiago"
	fte.age = 25
	fte.vacation = false
	fmt.Printf("%+v", fte)

	// Does not work because Go does not implement inheritance
	// fmt.Println(fte.GetName())
}
