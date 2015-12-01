package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var s string //declare a variable of type string

	// pass in the variable as a reference
	// this function will operate on the same reference that I declared in the calling function
	// any changes made in the function will be reflected here

	// The fmt.Scanln() function is used for simple input and for parsing a string
	// it automatically breaks up the input wherever it finds space charachters
	// if you want to simply collect user input, you should instead use a couple of packages
	// called bufio and os
	// fmt.Scanln(&s) 
	// fmt.Println(s)


	// create a reader object
	// (os.Stdin means that the reader is looking for standard user input (from the console))
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)


	// what if you want a number instead
	fmt.Print("Enter a number: ")
	// notice that the colon is removed.
	// this means you are assigning a value to a predeclared variable
	// := means declaration I assume
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}