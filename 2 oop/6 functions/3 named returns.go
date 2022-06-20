package main

import "fmt"

// declare values on the return declaration
func sum(n int) (double, triple, quad int) {
	double = 2 * n
	triple = 3 * n
	quad = 4 * n
	return
}

func main() {
	fmt.Println(sum(1))
}
