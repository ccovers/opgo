package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(delta int) {
			time.Sleep(1 * time.Second)
			fmt.Printf("hello %d\n", delta)
			//wg.Add(-delta)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("ok\n")
}
