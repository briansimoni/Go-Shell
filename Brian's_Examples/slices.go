/* From Wikipedia:
	This article is about the data structure operation. For other uses of slicing, see Slicing (disambiguation).
In computer programming, array slicing is an operation that extracts certain elements from an array and packages them as another array, possibly with different number of indices (or dimensions) and different index ranges. Two common examples are extracting a substring from a string of characters (e.g. "ell" from "hello"), and extracting a row (or a column) of a rectangular matrix to be used as a vector.

Depending on the programming language and context, the elements of the new array may be aliased to (i.e., share memory with) those of the original array.





A slice in Go is an abstraction layer that sits on top of an array
When you declare a slice, the runtime creates
1. the underlying array for you
2. allocates the required memory
3. and then returns the requested slice

unlike arrays they are resizable

*/


package main

import(
	"fmt"
	"sort"
)

func main(){
	var colors = [] string{"Red", "Green", "Blue"}//without a number it means it is a slice (not an array)
	fmt.Println(colors)

	// you can resize slices, use the append function
	colors = append(colors, "Purple")
	fmt.Println(colors)

	// remove the first item from the slice
	colors = append(colors[1:len(colors)]) // (len(colors) is the default)
	fmt.Println(colors)

	colors = append(colors[1:])
	fmt.Println(colors)

	colors = append(colors[:len(colors) -1])
	fmt.Println(colors)


	// declare a slice named numbers, make slice of ints, initial size of 5
	numbers := make([] int, 5, 5,)
	numbers[0] = 9
	numbers[1] = 8
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	// if you go over the cap it will just increase the size
	// it will not go line by line
	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)
}