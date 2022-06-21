package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		Buffered channel as Semaphore to limit
		the number of go routines executed at once.
	*/
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	// c := [][]
	// c := [goRoutine1][goRoutine2]
	// c := [goRoutine3][goRoutine2]

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Println("Started:", i)
	time.Sleep(1 * time.Second)
	fmt.Println("End:", i)
	<-c
}
