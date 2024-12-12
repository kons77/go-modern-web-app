package main

import (
	"fmt"
	"net/http"
)

// the basis for  any web application in go
func main() {
	// HandleFunc is listening for a request sent by a web browser
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	// enter localhost to a browser / 127.0.0.1:8080
	_ = http.ListenAndServe(":8080", nil)
}

// pointer is an address in memory where some value is stored
