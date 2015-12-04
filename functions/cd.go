package functions

import (
    "fmt"
    "path/filepath"
    "os"
)


func Cd(path, arg string, err error) string{
	path, err = filepath.Abs(arg)
	os.Chdir(path)
	if err != nil{
		fmt.Println(err)
	}
	return path
}