package main

import (
    "fmt"
    "os"
)

func main() {
    argLength := len(os.Args[1:])
    fmt.Printf("Number of Args is %d", argLength)
}

