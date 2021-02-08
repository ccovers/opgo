package main

import (
	"fmt"
	"sync"
	"time"
)

var MaxNum int = 1000000

func main() {
	tm := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	// channel 测试
	go func() {
		ch := make(chan int)
		go receiveChan(ch)
		for i := 0; i < MaxNum; i++ {
			ch <- i
		}
		close(ch)
		fmt.Printf("chan spend: %v\n", time.Since(tm))
		wg.Done()
	}()
	// mutex 测试
	go func() {
		var lock sync.Mutex
		m := make(map[int]int, MaxNum)
		for i := 0; i < MaxNum; i++ {
			lock.Lock()
			m[i] = i
			lock.Unlock()
		}
		fmt.Printf("mutex spend: %v\n", time.Since(tm))
		wg.Done()
	}()
	wg.Wait()
}

func receiveChan(ch <-chan int) {
	m := make(map[int]int, MaxNum)
	for {
		select {
		case k, ok := <-ch:
			if !ok {
				return
			}
			m[k] = k
		}
	}
}
