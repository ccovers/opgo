package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    ch := make(chan int, 0)
    var wg sync.WaitGroup

    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(x int) {
            defer wg.Done()
            ch <- x
        }(i)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for {
        select {
        case v, ok := <-ch:
            if !ok {
                goto OVER
            } else {
                fmt.Printf("v: %d\n", v)
            }
        case time.After(1 * time.Second):
            fmt.Println("===")
        default:
            time.Sleep(1 * time.Second)
        }
    }
OVER:
    fmt.Println("game over")
}
