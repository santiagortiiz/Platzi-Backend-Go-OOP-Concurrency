/*
Interfaces are like contracts that restrict the classes
who's implement them, to define all the methods that
the interface has.
*/
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id       int
	vacation bool
}

// class with the getMessage method (implicit interface)
type FullTimeEmployee struct {
	// Anonymous fields define the composition
	Person
	Employee
	endDate string
}

func (fte FullTimeEmployee) getMessage() string {
	return "I'm a full time employee"
}

// class with the getMessage method (implicit interface)
type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (te TemporaryEmployee) getMessage() string {
	return "I'm a temporary employee"
}

// Interface declaration
type PrintInfo interface {
	// Syntax:
	// method() return_type
	getMessage() string
}

// Interface implementation
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	fte := FullTimeEmployee{}
	te := TemporaryEmployee{}

	getMessage(fte)
	getMessage(te)
}
