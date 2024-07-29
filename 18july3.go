package main

import (
	"fmt"
	"net/http"
)

// Handler for the root route "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

// Handler for the "/hello" route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// Handler for the "/goodbye" route
func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye, World!")
}

func main() {
	// Register handlers for different routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
