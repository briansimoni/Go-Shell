package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"

	fmt.Println(str1) // prints out "An implicitly typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1) // print value, print type

	str2 := "An explicitly typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1)) // use the strings package to do stuff with strings

	fmt.Println(strings.Title(str1)) // first letters to upper case of each word

	firstValue := "hello"
	secondValue := "HELLO"
	fmt.Println("Equal?", (firstValue == secondValue)) // false
	fmt.Println("Equal Non-Case_Sensitive?", strings.EqualFold(firstValue, secondValue)) //true
	fmt.Println("Contains exp?", strings.Contains(str1, "exp")) // returns boolean
}