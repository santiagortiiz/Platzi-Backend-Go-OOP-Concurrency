package main

import "fmt"

func sum(values ...int) int {
	var total int = 0
	for _, v := range values {
		total += v
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2))
}
