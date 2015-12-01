package main

import(
	"fmt"
	"math/rand"
	"time"
)

// breaks are not necessary in go switch statements
// fallthrough statement is available if you want the 'fallthrough behavior'



func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	// fmt.Println("Day", dow)



	result := ""

	// like if statements, you can make variables only available 
	// to conditional logic (then available for garbage collection)

	switch dow := rand.Intn(6) + 1; dow { 
	case 1:
		result = "It's Sunday!"
	case 2:
		result = "It's Monday!"
	case 3:
		result = "It's Tuesday!"
	case 4:
		result = "It's Wednesday!"
	case 5:
		result = "It's Thursday!"
	default:
		result = "It's either Friday or Saturday"
	}

	fmt.Println("Day", dow, ",", result)

	x := -42
	switch {
	case x < 0:
		result = "Less than zero"
	case x == 0:
		result = "Equals zero"
	case x > 0:
		result = "Greater than 0"
	}

	fmt.Println(result)
}