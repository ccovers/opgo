package main

import (
	"fmt"
	"sync"
	"time"
)

type HanleParam struct {
	Data interface{} `comment:"待处理的数据，会传入回调函数中"`
	Num  int         `comment:"调用协程的次数"`
	Mx   *sync.Mutex `comment:"互斥锁"`
}

func handlePanic() func() {
	return func() {
		if r := recover(); r != nil {
			return
		}
	}
}

/*
* param:
* 	@ maxGoNum 协程最大数量
* 	@ timeout 超时时间
*	@ handle 回调处理方法
* 	@ param 待处理的参数
 */
func GoHandle(
	maxGoNum int, timeout time.Duration,
	handle func(*HanleParam, int), param *HanleParam,
) error {
	goPool := make(chan struct{}, maxGoNum) // 限制同时最多开启 maxGoNum 个协程
	signal := make(chan struct{})           // 通知请求完成
	defer close(goPool)
	defer close(signal)

	go func() {
		defer handlePanic()()
		for i := 0; i < param.Num; i++ {
			// 每个数据创建协程处理时向chanel中写数据，超过限制的协程序数量时，阻塞
			goPool <- struct{}{}

			go func(param *HanleParam, index int) {
				defer handlePanic()()
				// 处理数据
				handle(param, index)
				// 协程处理完毕，释放限制
				<-goPool
				// 处理完成通知
				signal <- struct{}{}
			}(param, i)
		}
	}()

	// 定时器
	if timeout <= 0 {
		timeout = 0x7FFFFFFF * time.Second
	}
	doTmer := time.After(timeout)
	// 数据处理通知，统计数量达成所有数据，则跳出循环
	for i := param.Num; i > 0; i-- {
		select {
		case <-signal:
		case <-doTmer:
			return fmt.Errorf("run timeout")
		}
	}
	return nil
}

var num int

func getNum() int {
	num += 1
	return num
}

func main() {
	data := []int{1, 2, 3, 4, 5}

	fmt.Println("=====")
	err := GoHandle(2, 10*time.Second,
		handle, &HanleParam{Data: data, Num: len(data), Mx: new(sync.Mutex)})
	if err != nil {
		fmt.Println("gohandle: ", err)
	}
	fmt.Println("=====...")
}

func handle(param *HanleParam, index int) {
	data, ok := param.Data.([]int)
	if !ok {
		return
	}
	time.Sleep(2 * time.Second)

	// 处理数据时，加锁
	param.Mx.Lock()
	fmt.Printf("goroutine[%d]: %d, num: %d\n", index, data[index], getNum())
	param.Mx.Unlock()
}
