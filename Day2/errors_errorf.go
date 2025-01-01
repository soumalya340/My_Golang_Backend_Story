package main

import "fmt"

func main() {

	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	// create an error using Efforf()
	error := fmt.Errorf("%d you are too young\nDon't come here 👋 with your age", age)

	if age < 18 {
		fmt.Println(error)
	} else {
		fmt.Println("You are welcome to the club, in your age  ", age, " , you age like a fine wine 🍷 ")
	}
}
