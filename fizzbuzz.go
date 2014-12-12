package main

import "fmt"

func main() {
    i := 0
    for i < 100 {
        i = i+1
        div := 0
        if i % 3 == 0 {
            fmt.Print("fizz")
            div = 1
        }
        if i % 5 == 0 {
            fmt.Print("buzz")
            div = 1
        }
        if div == 0 {
            fmt.Print(i)
        }
        fmt.Println()
    }
}