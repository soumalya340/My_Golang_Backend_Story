package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var number int
	fmt.Println("Enter a number for it's factorial:")
	fmt.Scanln(&number)
	fmt.Println(factorial(number))
}

