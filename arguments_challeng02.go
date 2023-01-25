package main

import (
    "fmt"
    "os"
)

func main() {

    argsLength := len(os.Args[1:])
    fmt.Printf("Arg length is %d\n", argsLength)
    for i, Args := range os.Args[1:] {
        fmt.Printf("Arg %d is %s \n", i+1,Args)
    }

}

