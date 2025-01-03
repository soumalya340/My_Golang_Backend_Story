package main

import (
	"fmt"
	engine "main.go/engine/server"
)

func main() {
	a := engine.Server{}
	a.Start()
	fmt.Println("Server is running on port 3333")
}
