package main

import (
    "fmt"
    "io/ioutil"
)

func main() {z
    files, _ := ioutil.ReadDir("./")
    for _, f := range files {
            fmt.Println(f.Name())
    }
}