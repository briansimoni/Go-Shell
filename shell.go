// runs the program, prints out the currenct directory + dollar sign
// waits for user input
// commands that work right now: ls, exit
// if command not found, print "No command (command) found"
package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
)

func main() {

	exit := false
	for exit == false {

		root, _ := filepath.Abs(".")
		fmt.Print(root,"$ ")

		input := ""
		fmt.Scanln(&input)

		if input == "ls" {

			files, _ := ioutil.ReadDir("./")
			for _, f := range files {
	    		fmt.Println(f.Name())
	    	}

		} else if input == "exit" {
			exit = true
			fmt.Println("Go shell exited")
		} else {
			fmt.Println("No command",input,"found")
		}

	}


}