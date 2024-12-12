package main

import (
	"webapp19/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main os the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	// enter localhost to a browser / 127.0.0.1:8080
	_ = http.ListenAndServe(portNumber, nil)
}
