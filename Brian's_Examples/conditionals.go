package main

import(
	"fmt"
)



func main() {

	// var x float64 = 0
	var result string

	// due to how the lexer works the else must be placed on the same line as the ending if bracket

	if x := 42; x < 0 { // you can instantiate variables in the condition (eligble for garbage collection)
		result = "Less than zero"
	}else if x == 0 {
		result = "equals zero"
		
	} else {
		result = "Greater than or equal to zero"
	}

	fmt.Println("Result:", result)
/*	fmt.Println("Value of x:", x)*/
}