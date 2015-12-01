package main

import(
	"fmt"
)

func main() {
	// all items in the array must be of the same type
	var colors [3] string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5] int{1,2,3,4,5} // array literals

	fmt.Println(numbers)

	fmt.Println("Number of colors:",len(colors)) // len is the same function you use to get the length of  a string
	fmt.Println("Number of numbers:",len(numbers))

	/* An array is an object, just like all objects in Go, if you pass in an object into a function
	a copy will be made of the array rather than a reference but storing data in memory is just about
	all you can do with simple arrays. You can't easily sort them and you can't add or remove items
	at runtime. For those features and more you should package your ordered data in slices instead
	of arrays
	*/
}