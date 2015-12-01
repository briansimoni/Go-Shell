/*

Go's defer statement does exactly what it says.
It defers execution of a block of code until everything else in the current functionis finished

Each time you call the defer statement, it adds an instruction to a stack
Handled in LIFO of course

How can you use deferred statements?

When you open a file or a database connection
You can defer the closing of the connection until everything is complete

*/

package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Close the file")
	fmt.Println("Open the file")

	myFunc()

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("Undeferred statement")

	x := 1000
	defer fmt.Println("Value of x:",x) // x is evaluated and value is saved at the moment of defering (this prints out 1000)
	x++
	fmt.Println("Value of x without defering it",x) 

}

func myFunc() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}