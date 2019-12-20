package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := []int{1, 2, 3, 4, 5}

	var mx sync.Mutex
	queue := make(chan struct{}, 5) // 限制同时最多开启5个协程
	signal := make(chan struct{})   // 通知请求完成
	defer close(queue)
	defer close(signal)

	fmt.Println("=====\n")
	for _, v := range data {
		// 每个数据创建协程处理时向chanel中写数据，超过限制的协程序数量时，阻塞
		queue <- struct{}{}

		go func(v int) {
			fmt.Printf("%d \n", v)

			// 处理数据时，加锁
			mx.Lock()
			time.Sleep(1 * time.Second)
			fmt.Println(v, ":", getNum())
			fmt.Printf("%d ...\n", v)
			mx.Unlock()

			// 协程处理完毕，释放限制
			<-queue
			// 处理完成通知
			signal <- struct{}{}
		}(v)
	}

	// 数据处理通知，统计数量达成所有数据，则跳出循环
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
