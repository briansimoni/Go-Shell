package main

import (
	"fmt"
	"stringutil"
)


// this is to demonstrate importing other packages/libraries or something

func main() {

	n1, l1 := stringutil.FullName("Poop", "Dood")
	fmt.Printf("FullName: %v, number of chars: %v\n", n1, l1)

	n1, l1 = stringutil.FullNameNakedReturn("Butt", "Face")
	fmt.Printf("FullName: %v, number of chars: %v\n", n1, l1)
	
}