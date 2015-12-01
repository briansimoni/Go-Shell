package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."

	aNumber := 42
	isTrue := true

	// go functions can return more than one value similtaneously.
	// stringLength, err := fmt.Println(str1, str2, str3)
	stringLength, _ := fmt.Println(str1, str2, str3) // use _ to throw away the variable

	// if err == nil {
		fmt.Println("String length:", stringLength)
	// }

		// %v means output the value of a variable
		// you need to pass in a variable for each %v
		// \n is for new line
		fmt.Printf("value of aNumber: %v\n", aNumber)

		fmt.Printf("value of isTrue: %v\n", isTrue)

		// to 2 decimal places. No implicit type conversion. Convert to float
		fmt.Printf("Value of aNumber as a float: %.2f\n", float64(aNumber))


		// find out the types of each variable
		fmt.Printf("Data types: %T, %T, %T, %T, %T,\n", str1, str2, str3, aNumber, isTrue)

		// Sprintf means to return a string	
		myString:= fmt.Sprintf("Data types as var: %T, %T, %T, %T, %T,\n", 
			str1, str2, str3, aNumber, isTrue)

		fmt.Println(myString)
}