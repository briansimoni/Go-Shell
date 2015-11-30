// runs the program, prints out the currenct directory + dollar sign
// waits for user input
// if command not found, print "No command (command) found"
//
//
// ****** commands that work right now: ls, exit, cd ******

package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "os"
    "bufio"
    "strings"
)

func main() {

	path, err := filepath.Abs("/")
	os.Chdir(path)

	exit := false
	for exit == false {

		
		fmt.Print(path ,"$ ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		slice := strings.Split(input, " ")

		fmt.Println(slice)

		command := slice[0]
		var arg string

		if len(slice) > 1{
			arg = slice[1]
		}


		if command == "dir" {

			files, _ := ioutil.ReadDir(path)
			for _, f := range files {
	    		fmt.Println(f.Name())
	    	}

		} else if command == "cd" {
			path, err = filepath.Abs(arg)
			os.Chdir(path)
			if err != nil{
				fmt.Println(err)
			}
		} else if command == "exit" {
			exit = true
			fmt.Println("Go shell exited")
		} else {
			fmt.Println("No command",command,"found")
		}

	}


}