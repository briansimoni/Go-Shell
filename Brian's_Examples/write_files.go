package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main() {
	content := "Hello from Go!"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with te file of %v charachters\n", ln)

	// you can do this if you need to create a binary file
	bytes := []byte(content)
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)
}


func checkError(err error) {
	if err != nil{
		panic(err) //panic function calls the application to stop and display error
	}

}