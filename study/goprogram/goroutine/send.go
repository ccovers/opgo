package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// GOMAXPROCS 确定执行代码的OS线程数量
func main() {
	fmt.Println(runtime.GOMAXPROCS(1))
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	var wg sync.WaitGroup
	cnt := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			ch1 <- struct{}{}
			cnt += 1
			_, ok := <-ch2
			if !ok {
				close(ch1)
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		ta := time.After(1 * time.Second)
		defer wg.Done()
		for {
			select {
			case _, ok := <-ch1:
				if !ok {
					close(ch2)
					return
				}
				ch2 <- struct{}{}
			case <-ta:
				close(ch2)
				return
			}
		}
	}()
	wg.Wait()
	fmt.Println("send cnt:", cnt)
}
