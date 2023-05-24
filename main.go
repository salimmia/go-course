package main

import (
	"errors"
	"fmt"
	"net/http"
	// "os"
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

func Divide(w http.ResponseWriter, r *http.Request){
	var x, y float32;

	x = 6;
	y = 0;

	result, err := divideValues(x, y);

	if err != nil{
		fmt.Fprintf(w, fmt.Sprint(err));
		return
		// os.Exit(1);
	}
	fmt.Fprintf(w, fmt.Sprintf("The result of %v/%v is %v", x, y, result));
}

func divideValues(x, y float32) (float32, error){
	if y == 0{
		err := errors.New("cannot divide by zero");

		return 0, err;
	}

	result := x / y;

	return result, nil;
}

// main is the main application function
func main(){
	http.HandleFunc("/home", Home);

	http.HandleFunc("/about", About);

	http.HandleFunc("/divide", Divide);

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber));

	_ = http.ListenAndServe(portNumber, nil);
}