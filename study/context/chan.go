package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Printf("监控退出\n")
				return
			default:
				fmt.Printf("goroutine监控中\n")
				time.Sleep(2 * time.Second)
				fmt.Printf("goroutine监控中...\n")
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Printf("通知监控停止\n")
	stop <- true
	time.Sleep(2 * time.Second)
}
