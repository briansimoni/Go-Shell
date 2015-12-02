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
            if err != nil {}

		case "dir":
			functions.Dir(path)

		case "quit":
			exit = true
			fmt.Println("Go shell exited")

		default:
			fmt.Println("No command", command, "found")
		}

	}

}
