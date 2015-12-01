package main


/* Go does not have a class try catch statement
an error in go is an instance of an interface that deffines a single method, named error,
and that method returns a string and that string is the error message */

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil{
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}


	myError := errors.New("My error string")
	fmt.Println(myError)

	attendance := map[string]bool{
		"Ann": true,
		"Mike": true}
	attended, ok := attendance["Mike"]

	if ok {
		fmt.Println("Mike attended?", attended)
	} else {
		fmt.Println("No info for mike")
	}
}
