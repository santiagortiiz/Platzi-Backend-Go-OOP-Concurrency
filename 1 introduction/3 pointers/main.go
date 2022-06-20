package main

import "fmt"

func main() {
	g := 25
	fmt.Println("g: ", g)
	h := &g
	fmt.Println("- reference &g: ", h)
	fmt.Println("- value of the pointer *h: ", *h)
}
