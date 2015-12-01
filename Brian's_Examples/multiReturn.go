package main

import "fmt"


// since there is no structured exception handling in go
// multi return functions are essential

func main() {

	n1, l1 := FullName("Poop", "Dood")
	fmt.Printf("FullName: %v, number of chars: %v\n", n1, l1)

	n1, l1 = FullNameNakedReturn("Butt", "Face")
	fmt.Printf("FullName: %v, number of chars: %v\n", n1, l1)
	
}

func FullName(f, l string) (string, int){ // a convention in go is you don't use the work 'get'
	full := f + " " + l
	length := len(full)
	return full, length
}


// called a naked return because when you use the return statement you do not have to specify
// the order of things you are returning up in the parantheses
func FullNameNakedReturn(f, l string) (full string, length int){
	full = f + " " + l
	length = len(full)
	return
}