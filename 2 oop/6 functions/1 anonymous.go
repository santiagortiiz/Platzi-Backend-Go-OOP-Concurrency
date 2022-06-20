package main

import (
	"fmt"
	"time"
)

func main() {
	// Doubles a number with an anonymous function
	x := 5
	y := func(n int) int {
		return x * n
	}(2)
	fmt.Println("x:", x, "| 2x", y)

	// Run a goroutine with an anonymous function
	c := make(chan int)
	go func() {
		fmt.Println("Long process")
		time.Sleep(1 * time.Second)
		fmt.Println("End Long process")
		c <- 1
	}()
	println(<-c)
}
