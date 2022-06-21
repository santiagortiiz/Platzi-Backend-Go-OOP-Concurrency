/*
An unbuffered channel is used to perform synchronous communication between goroutines
while a buffered channel is used for perform asynchronous communication.
An unbuffered channel provides a guarantee that an exchange between two goroutines
is performed at the instant the send and receive take place.

Unbuffered channel:
Espera una función o una rutina para recibir el mensaje,
es bloqueada por ser llamada en la misma función.

Buffered channel: Se puede llamar de manera inmediata,
en el siguiente ejemplo 2 es el numero de canales que pueden ser usados.
*/

package main

import "fmt"

func main() {
	// Unbuffered channels
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
