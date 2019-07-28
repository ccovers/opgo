package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chs := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			chs <- x
		}(i)
	}

	go func() {
		wg.Wait()
		close(chs)
	}()

	for x := range chs {
		fmt.Println(x)
	}
	fmt.Println("over")
}
