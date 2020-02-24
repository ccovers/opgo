package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	watch := func(ctx context.Context, name string) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("%s: 监控退出\n", name)
				return
			default:
				fmt.Printf("%s: goroutine监控中\n", name)
				time.Sleep(1 * time.Second)
				// fmt.Printf("%s: goroutine监控中...\n", name)
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		key := fmt.Sprintf("key_%d", i)
		val := fmt.Sprintf("[监控_%d]", i)
		vctx := context.WithValue(ctx, key, val)
		go watch(ctx, val)
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("通知监控停止\n")
	cancel()
	wg.Wait()
	fmt.Printf("程序退出\n")
}
