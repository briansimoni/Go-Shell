// Go contains pretty much all the same Math operations that Java and C have

package main

import (
	"fmt"
	"math/big" //subdirectory of the math package
	"math"
)

func main() {
	i1, i2, i3 := 12, 45, 68 // you can declare multiple variables on one line
	intSum := i1 + i2 + i3
	fmt.Println("integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("float sum: ", floatSum) // same typical probablems with floating point

	var b1, b2, b3, bigSum big.Float

	b1.SetFloat64(23.5)
	b2.SetFloat64(65.1)
	b3.SetFloat64(76.3)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3) // the ampersand indicates that they are pointers
	fmt.Printf("BigSum - %.10g\n", &bigSum)

	// without the pointer syntax you see the object
	// thus you will see all of the fields and shit

	circleRadius := 15.5
	circumference := circleRadius * math.Pi
	fmt.Printf("circumference: %.2f\n", circumference) // %.2f print the value to 2 decimal places
}