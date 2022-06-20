package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

// Define a Struct in a function (Constructor 4)
func NewEmployee(id int, name string, vacation bool) *Employee {
	// Returns a reference to the Struct instead of a copy
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Constructor 1: default values
	e1 := Employee{}
	fmt.Printf("%+v\n", e1)

	// Constructor 2: using "New"
	e2 := new(Employee)
	fmt.Printf("%+v", e2) // This way returns the reference to the employee
	fmt.Printf("%+v\n", *e2)

	// Constructor 3: defining properties values
	e3 := Employee{
		id:       1,
		name:     "Santiago",
		vacation: false,
	}
	fmt.Printf("%+v\n", e3)

	// Constructor 4: functions
	e4 := NewEmployee(1, "New Employee", false)
	fmt.Printf("%+v\n", e4)
}
