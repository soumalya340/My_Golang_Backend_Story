package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	fmt.Println("Table of ", number, " is: ")
	var count = 1
	var result int
	for result < number*10 {
		result = number * count
		fmt.Printf("%d * %d = %d\n", number, count, result)
		count++

	}

}
