package main

import (
	"github.com/donvito/hellomod"
	// alias "package"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Santiago, ")
}
