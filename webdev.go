package main

import (
	"fmt"
	"net/http"
)

// Handler function for the root route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Web Development!")
}

func main() {
	// Define routes and handlers
	http.HandleFunc("/", helloHandler)

	// Start the web server
	port := 8080
	fmt.Printf("Server is running on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
