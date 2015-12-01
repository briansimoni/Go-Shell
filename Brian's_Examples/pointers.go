package main

import "fmt"

func main(){
	var p *int

	if p != nil{
		fmt.Println("Value of p: ", p)
	} else {
		fmt.Println("p is nil")
	}


	var v int = 42
	p = &v //ampersand means to connect the pointer to this variable

	if p != nil{
		fmt.Println("Value of p: ", *p) //reference the value throug the pointer with the * operator
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1", *pointer1)
	fmt.Println("Value1:", value1) // again * is used to reference the pointer itself, and & is how you connect the value to the pointer

}