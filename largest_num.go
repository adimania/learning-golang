package main

import "fmt"

func main() {
    x := [10]int{3,5,7,4,6,1,8,6,9,0}
    large := 0
    for _, current := range x {
       if current > large {
           large = current
       }
    }
    fmt.Println(large)
}