package main

import "fmt"

func counterMaker()  int {
	count := 0
	count++
	return count
}

func main() {
	// counter := counterMaker()

	fmt.Println(counterMaker()) // prints 1
	fmt.Println(counterMaker()) // prints 2
	fmt.Println(counterMaker()) // prints 3
}