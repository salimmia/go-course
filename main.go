package main

import (
	"fmt"
	"net/http"
	// "os"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/home", Home)

	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
