/* Defining and calling functions */

// go is organized with packages, and packages have functions
// your own application has its own package, always named main
// and it also has the main function which is called automatically by the run time
// as the application starts up. You can define yourr own custom functions and you
// can then organize them into their own packages



package main

import "fmt"

func main() {
	doSomething()
	sum := addValues(1,2)
	sum = addAllValues(12,13,14)
	fmt.Println(sum)


}

// you cannot have two dfirrent versin of the function with
// different numbers of arguments or types of arguments
// i.e. no overloading functions
// since there is no inheritance, I assume there is no overriding functions either
func doSomething() {
	fmt.Println("Doing Someting")
}

// if you have arguments of the same type, you can just pass in a list
// and declare the type only once
func addValues(value1, value2 int) int { // the thing after the parantheses is the return type
	return value1 + value2
}


func addAllValues(values ...int) int{ // this creates a slice with an arbitrary number of ints
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values) // %T gets the type
	return sum
}

// remember that all of these functions are private because of the lower case charachter