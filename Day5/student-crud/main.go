package main

import (
	"fmt"
	"log"
	"net/http"
	"student-crud/handlers"
)

func main() {
	// Initialize routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/students", handlers.StudentsHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start server
	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
