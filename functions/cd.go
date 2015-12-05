package functions

import (
"fmt"
"path/filepath"
"os"
"io/ioutil"
)



func Cd(arg string) string {

	path, _ := filepath.Abs(arg)

	_, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("error:",err)
		return os.Getenv("PWD")
	} else {
		os.Chdir(path)
		return path
	}

}