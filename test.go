package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mx sync.Mutex
	data := []int{1, 2, 3, 4, 5}
	queue := make(chan struct{}, 2)
	signal := make(chan struct{})
	defer close(queue)
	defer close(signal)

	fmt.Println("=====\n")
	for _, v := range data {
		queue <- struct{}{}
		go func(v int) {
			fmt.Printf("%d \n", v)
			time.Sleep(1 * time.Second)
			mx.Lock()
			fmt.Println(v, ":", getNum())
			mx.Unlock()
			fmt.Printf("%d ...\n", v)
			<-queue
			signal <- struct{}{}
		}(v)
	}

	cnt := 0
	for range signal {
		cnt += 1
		if cnt == len(data) {
			break
		}
	}
	fmt.Println("=====...\n")
}

var num int

func getNum() int {
	num += 1
	return num
}
