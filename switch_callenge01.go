/* RZFeeser | Alta3 Research
   switch - case and default  */

package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

    var gove string = runtime.Version()
    fmt.Println(gove)

           // init gover                
    switch {
    case strings.Contains(gove, "go1.18"):
        fmt.Println("you are using the latest version of GoLang")
    case strings.Contains(gove, "go1.19"):
        fmt.Println("You ar using the uptodate version of GoLang")
    default:                       // in all other cases run the code below
        fmt.Println("I cannot make a recommendation.")
    }
}

