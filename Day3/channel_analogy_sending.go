package main

import (
    "fmt"
    "time"
)

// Function for the gossiping friend (First sender)
func gossipingFriend(ch1 chan string) {
    fmt.Println("Still gossiping...")
    time.Sleep(2 * time.Second)  // Gossiping time!
    ch1 <- "Done gossiping, here's your pencil!"
}

// Function for the middle student (You - both receiver and sender)
func middleStudent(ch1, ch2 chan string) {
    fmt.Println("Asked for my pencil...")
    pencil := <-ch1  // Blocked until gossiping friend gives pencil
    fmt.Println(pencil)
    fmt.Println("Got it! Now passing to brushing friend...")
    
    ch2 <- "Hey brushing friend, here's the pencil!"  // Blocked until brushing friend receives
}

// Function for the brushing friend (Final receiver)
func brushingFriend(ch2 chan string) {
    fmt.Println("Brushing teeth... 🦷")
    time.Sleep(3 * time.Second)  // Brushing time!
    message := <-ch2
    fmt.Println(message)
    fmt.Println("Thanks! Done brushing! 😁")
}

func main() {
    ch1 := make(chan string)  // Channel between gossiping friend and you
    ch2 := make(chan string)  // Channel between you and brushing friend
    
    // Start both friends in goroutines
    go gossipingFriend(ch1)
    go brushingFriend(ch2)
    
    // You're in the middle, receiving and sending
    middleStudent(ch1, ch2)
}