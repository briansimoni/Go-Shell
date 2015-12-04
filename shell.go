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
//    "io"
    "io/ioutil"
)

func main() {

    // define and set environment variables
    path := os.Getenv("PWD")
    parent := "Go-Shell"
    shell := searchPATHforProgram(parent)
    
    os.Setenv("PARENT", parent)
    os.Setenv("SHELL", shell)
    
    if (len(os.Args) == 1){         // start the REPL in interactive mode
        repl("", path)
    } else {                        // start the REPL to read each of the files
        for _, f := range os.Args[1:] {
            repl(f, path)
        }
    }


} // end main



///////////////////////////////////////////////////////////////////////////
// read, eval, print, loop
/////////////////////////////////////////////////////////////////////////// 
func repl (filename, path string) error {  
    
    var commandlinesource *bufio.Scanner 
    
    if (filename == "") {
        commandlinesource = bufio.NewScanner(os.Stdin)
    } else {
        commandfile, _ := os.Open(filename)
        defer commandfile.Close()
        
        commandlinesource = bufio.NewScanner(commandfile)
    }
      
    exit := false
    for exit == false { // essentially a while loop
        
        // Display a prompt
        fmt.Print(path, "$ ")
        

        // read a line of input
        var input string
        if( commandlinesource.Scan() ){
            input = commandlinesource.Text()
        } else {
            exit = true
        }
        input = strings.TrimSpace(input)
        tokens := strings.Split(input, " ")

        // set the command to the first token and its arguments to the rest
        command := tokens[0]
        var args []string
        if len(tokens) > 1 {
            args = tokens[1:]
        }
        
        // Advance the display writer to the next line
        fmt.Println()

        // execute commands
        switch command {
            
            case "":
                // do nothing
                
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
                s := strings.Join(args, " ")
                fmt.Println(s)
            case "pause":
                fmt.Print("Press 'Enter' to continue...")
                bufio.NewReader(os.Stdin).ReadBytes('\n')
                
            case "quit":
                exit = true
                fmt.Println("Go shell exited")

            case "help":
            	fmt.Println("cd - change the current directory");
            	fmt.Println("clr - clear the screen");
            	fmt.Println("dir - list the contents of the directory");
            	fmt.Println("environ - list all environment strings");
            	fmt.Println("echo [comment] - prints the comment to the screen");
            	fmt.Println("help - displays the user manual");
            	fmt.Println("pause - pauses the shell until enter is pressed");
            	fmt.Println("quit - quits the shell");

            default:
                err := invokeProgram (command, args)
                if (err != nil) {
                    fmt.Println("No command", command, "found")
                }
            
        } //end command switch

    } // end command loop
/////////////////////////////////////////////////////////////////////////////

    return nil
}





/*
 * invokeProgram - runs a program with its arguments as a new process
 * Input: the name of the program to run and a list of its arguments
 * Output: an error on whether or not the new process started
 */
func invokeProgram (programName string, args []string) error {
    
    fullpath := searchPATHforProgram(programName)

    if (fullpath == "") {
//         errorstr := strings.Join(["Command", programName, "not found"])
        errorstr := "Command" + programName + "not found"
        return errors.New(errorstr)
    }
    
    // os.StartProcess used instead of exec.Command because it cleanly accepts
    // an arbitrary slice of strings for a program and its arguments.
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
