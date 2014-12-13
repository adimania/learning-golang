package main

import "fmt"

func main() {
    var num int
    fmt.Scanf("%d", &num)
    fmt.Println(fib(num))
}

func fib(n int) int {
    if n < 2 {
        return n
    } else {
        return fib(n-1) + fib(n-2)
    }
}