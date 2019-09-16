package main

import (
    "fmt"
)

func main() {
    ch := make(chan int, 1)
    for i := 0; i < 10; i++ {
        fmt.Println("===: ", i)
        select {
        case x := <-ch:
            fmt.Println("x: ", x)
        case ch <- i:
            fmt.Println("i: ", i)
        }
    }
}
