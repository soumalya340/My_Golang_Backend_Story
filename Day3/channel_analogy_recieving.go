package main

import (
	"fmt"
	"time"
)

// Function for the gossiping friend
func gossipingFriend(ch chan string) {
	fmt.Println("Still gossiping...")
	time.Sleep(2 * time.Second) // Gossiping time!
	ch <- "Done gossiping, here's your pencil!"
}

// Function for the student waiting for pencil
func waitingStudent(ch chan string) {
	fmt.Println("Asked for my pencil...")
	pencil := <-ch // Blocked here until friend stops gossiping
	fmt.Println(pencil)
	fmt.Println("Finally: 😝")
}

func main() {
	ch := make(chan string)

	// Start the gossiping friend in a goroutine
	go gossipingFriend(ch)

	// Wait for the pencil
	waitingStudent(ch)
}
