package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path", root)

	err := filepath.Walk(root, processPath)
	if err != nil{
		fmt.Println("error:",err)
	}

}

// this function is super useful apparently
// check out the documenation on the Go website for more information
// once you are in this function you can basically do anything you
// need to do with a directory or files
func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory:",path)
		} else {
			fmt.Println("File:",path)
		}
	}

	return nil
}