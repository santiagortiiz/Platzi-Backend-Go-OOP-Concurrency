package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// The wg must be passed as reference, otherwise it will be passed as a copy!
		go doSomething(i, &wg)
	}

	// Block the program until the waitgroup counter is 0
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	// Decrease the WaitGroup counter by 1
	defer wg.Done()

	fmt.Println("Started\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
