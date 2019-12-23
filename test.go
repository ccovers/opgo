package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type HanleParam struct {
	Data interface{}
	Num  int
	Mx   *sync.Mutex
}

func GoHandle(goNum int, handle func(*HanleParam, int), param *HanleParam) {
	queue := make(chan struct{}, goNum) // 限制同时最多开启 goNum 个协程
	signal := make(chan struct{})       // 通知请求完成
	defer close(queue)
	defer close(signal)

	for i := 0; i < param.Num; i++ {
		// 每个数据创建协程处理时向chanel中写数据，超过限制的协程序数量时，阻塞
		queue <- struct{}{}

		go func(param *HanleParam, index int) {
			// 处理数据
			handle(param, index)
			// 协程处理完毕，释放限制
			<-queue
			// 处理完成通知
			signal <- struct{}{}
		}(param, i)
	}

	// 数据处理通知，统计数量达成所有数据，则跳出循环
	cnt := 0
	if cnt < param.Num {
		for range signal {
			cnt += 1
			if cnt == param.Num {
				break
			}
		}
	}
}

var num int

func getNum() int {
	num += 1
	return num
}

func main() {
	fmt.Println(strings.Contains("asdjklll", "djo"))
	return

	data := []int{1, 2, 3, 4, 5}

	fmt.Println("=====\n")
	GoHandle(2, handle, &HanleParam{Data: data, Num: len(data), Mx: new(sync.Mutex)})
	fmt.Println("=====...\n")
}

func handle(param *HanleParam, index int) {
	data, ok := param.Data.([]int)
	if !ok {
		return
	}
	time.Sleep(1 * time.Second)

	// 处理数据时，加锁
	param.Mx.Lock()
	fmt.Println(data[index], ":", getNum())
	param.Mx.Unlock()
}
