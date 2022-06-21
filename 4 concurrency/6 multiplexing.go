package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	duration1 := 4 * time.Second
	duration2 := 2 * time.Second

	go DoSomething(duration1, c1, 1)
	go DoSomething(duration2, c2, 2)

	/*
		The value of the channel 2 arrives first, but it is shown at last,
		because the program is locked by the order of the Println.
	*/
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	/*
		To avoid this behavior, channels can be consumed
		without an specific order using a switch case statement.
	*/
	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)

		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
