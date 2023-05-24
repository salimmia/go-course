package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080";

// Home is the Home page Handler
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the Home page\nand My Name is Muhammad Selim");
}

// About is the about page Handler
func About(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "This is the about page and the value of 2 + 2 = %d", addValue(2, 2));
}


// addValue add two integers and return the sum
func addValue(x, y int) int{
	return x + y;
}

// main is the main application function
func main(){
	http.HandleFunc("/home", Home);

	http.HandleFunc("/about", About);

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber));

	_ = http.ListenAndServe(portNumber, nil);
}