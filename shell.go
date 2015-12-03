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
    "errors"
	"fmt"
	"os"
    "os/exec"
// 	"path/filepath"
	"strings"
    "io/ioutil"
)

func main() {

//     shellName := "Go-Shell"
  
    path := os.Getenv("PWD")
     
    // set the SHELL environment variable
    
    
//     os.Setenv("SHELL", shellName )
    
    
    
    

	exit := false
	for exit == false { // essentially a while loop

		fmt.Print(path, "$ ")

		// get user input
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		slice := strings.Split(input, " ")

		command := slice[0]
		var args []string

		if len(slice) > 1 {
			args = slice[1:]
		}

		// execute commands
		switch command {

		case "cd":
            if len(args) == 0 {
                fmt.Println(path)
            } else {
                path = functions.Cd(path, args[0], nil)
            }
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
        case "echo":
        	fmt.Println(args[0])
        case "pause":
        	fmt.Print("Press 'Enter' to continue...")
  			bufio.NewReader(os.Stdin).ReadBytes('\n')
            
		case "quit":
			exit = true
			fmt.Println("Go shell exited")

		default:
            err := invokeProgram (command, args)
            if (err != nil) {
                fmt.Println("No command", command, "found")
            }
            
		} //end command switch

	} // end command loop

} // end main

func invokeProgram (programName string, args []string) error {
    fullpath := searchPATHforProgram(programName)

    if (fullpath == "") {
//         errorstr := strings.Join(["Command", programName, "not found"])
        errorstr := "Command" + programName + "not found"
        return errors.New(errorstr)
    }
    
//     fmt.Println(fullpath)
    
//     cmd := exec.Command( fullpath, args )
//     cmd.Stdout = os.Stdout
//     cmd.Run()
    var attribs os.ProcAttr
    attribs.Files = []*os.File{ nil, os.Stdout, os.Stderr }
    
    // Need to include program name as ARGV[0]
    args = append( []string{programName}, args... )
    
    // Start the process
    proc, err := os.StartProcess(fullpath, args, &attribs)
    
    if (err != nil) {
        return errors.New("Error when starting " + programName)
    }
    proc.Wait()

    return nil
}


/*
 * searchPATHforProgram - returns the full path to a program's executable file
 * Input: the name of a program to search for
 * Output: the full path (including program filename) if successful
 *         else, an empty string.
 */
func searchPATHforProgram (programName string) string {
    
    programPath := ""
    
    // Get all the possible directories for the program to hide in
    executableDirs := strings.Split(os.Getenv("PATH"), ":")
    executableDirs = append(executableDirs, os.Getenv("GOPATH") )
    
    
    // Search through each directory to find a matching filename
    isFound := false
    for _, dir := range executableDirs {
        files, _ := ioutil.ReadDir(dir)
        
        for _, f := range files {
            if( f.Name() == programName ) {
                programPath = dir + "/" + programName 
                isFound = true
                break
            }
        }
        if (isFound) {
            break
        }
    }

    return programPath
}
