package functions

import (
    "fmt"
    "io/ioutil"
)


func Dir(path string){
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}