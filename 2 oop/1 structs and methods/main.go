package main

import "fmt"

// "class" definition
type Employee struct {
	// properties
	id   int
	name string
}

// "class" methods with receiver functions
// syntax: func (var_name *class) func_name(args) returns
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	// create an instance of the class
	e := Employee{}
	fmt.Printf("%+v", e)
	fmt.Println("")

	// modify it's properties
	e.id = 1
	e.name = "santiago"
	fmt.Printf("%+v", e)
	fmt.Println("")

	// call "class" methods
	e.SetId(5)
	e.SetName("Santiago Ortiz")
	fmt.Printf("%+v", e)
	fmt.Println("")

	fmt.Println("id:", e.GetId(), "name:", e.GetName())

}
