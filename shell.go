// runs the program, prints out the currenct directory + dollar sign
// waits for user input
// if command not found, print "No command (command) found"
//
//
// ****** commands that work right now: ls, exit, cd ******

package main

import (
	"Go-Shell/functions"
	"bufio"
	"fmt"
	"os"
    "os/exec"
	"path/filepath"
	"strings"
	//"io/ioutil"
	"io"
	"bytes"
)

func main() {

	path, err := filepath.Abs("/")
	os.Chdir(path)

	exit := false
	for exit == false { // essentially a while loop

		fmt.Print(path, "$ ")

		// get user input
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		slice := strings.Split(input, " ")

		command := slice[0]
		var arg string

		if len(slice) > 1 {
			arg = slice[1]
		}

		// execute commands
		switch command {

		case "cd":
			path = functions.Cd(path, arg, err)
        
        case "clr":
            cmd := exec.Command("clear")
            cmd.Stdout = os.Stdout
            cmd.Run()

		case "dir":
			functions.Dir(path)
        
        case "environ":
            env_vars := os.Environ()
            for _, env_str := range env_vars {
                fmt.Println(env_str)
            }
            
		case "quit":
			exit = true
			fmt.Println("Go shell exited")

		case "help":
			

		default:
			fmt.Println("No command", command, "found")
		}

	}

}