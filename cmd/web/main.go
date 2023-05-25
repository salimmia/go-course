package main

import (
	"fmt"
	"net/http"

	"github.com/salimmia/go-course/pkg/handlers"
	// "os"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/home", handlers.Home)

	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
