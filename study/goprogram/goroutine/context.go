package main

import (
    "context"
    "fmt"
    "sync"
)

type MyContext struct {
    Context context.Context
    Done    chan struct{}
    Mutex   sync.Mutex
    Map
}

func main() {
    myContext := MyContext{
        Context: context.Background(),
        Done:    make(chan struct{}),
        Mutex:   sync.Mutex,
    }

    c := context.WithDeadline(parent, d)
}
