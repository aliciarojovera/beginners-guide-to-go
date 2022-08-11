package main

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	// random
	value := rand.Intn(11)
	fmt.Printf("Value Calculated: %d\n", value)
	values <- value
}

func main() {
	values := make(chan int)
	go CalculateValue(values)
	go CalculateValue(values)
 	// <- block until something is received
	value := <-values
	value2 := <-values
	fmt.Println(value)
	fmt.Println(value2)
}
