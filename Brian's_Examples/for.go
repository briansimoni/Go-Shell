package main

import "fmt"

// There are no while loops in go!
// you can use extended versions of for

func main() {
	sum := 1
	fmt.Println("Sum:",sum)

	colors := []string{"Red","Green","Blue"} // remember without specifying the length, it is a slice
	fmt.Println(colors)

	sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}

	fmt.Println("Sum:",sum) // 5050

	for i := 0; i <len(colors); i++ {
		fmt.Println(colors[i])
	}

	// set the value of i to the current index as you loop through the slice
	// this is cool
	for i := range colors {
		fmt.Println(colors[i])
	}


	// break, continue, goto are all supported
	sum = 1
	for sum <= 1000 { // essentially a while loop here
		sum += sum
		fmt.Println("Sum:",sum) // 1024 is the final result


		if sum > 200 {
			goto endofprogram
		}

	}

	endofprogram : fmt.Println("end of program") // this is a label

}