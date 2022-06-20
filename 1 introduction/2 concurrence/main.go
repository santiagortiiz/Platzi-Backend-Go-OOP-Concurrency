package main

import (
	"fmt"
	"time"
)

func doSomething(c chan<- int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}

func main() {
	fmt.Println("End")

	// Channels
	c := make(chan int)
	go doSomething(c)
	fmt.Println(<-c)
}
