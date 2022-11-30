package main

import "fmt"

func main() {
    x := os.Args[1:]
    for _, v := range x {
        if x == '01' || x == "galaxy" || x == "galaxy 01" {
            fmt.Println("Alert!!!")
            return
        }
    }
}
