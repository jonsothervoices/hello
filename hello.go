package main

import (
	"fmt"
	"hello/mathexercises"
)

func main() {
	fmt.Println("Hello, world!")
	// stringexercises.FizzBuzz()
	fmt.Println(mathexercises.SieveOfEratosthenes(1000000))
}
