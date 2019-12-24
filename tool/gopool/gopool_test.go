package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var num int

func getNum() int {
	num += 1
	return num
}

func handle(param *HanleParam, index int) {
	_, ok := param.Data.([]int)
	if !ok {
		return
	}
	time.Sleep(1 * time.Second)

	// 处理数据时，加锁
	param.Mx.Lock()
	fmt.Printf("num: %d\n", getNum())
	param.Mx.Unlock()
}

func ExampleGopool() {
	data := []int{1, 2, 3, 4, 5}

	err := Handle(2, 0*time.Second,
		handle, &HanleParam{Data: data, Num: len(data), Mx: new(sync.Mutex)})
	if err != nil {
		fmt.Println("gohandle: ", err)
	}
	// Output:
	// num: 1
	// num: 2
	// num: 3
	// num: 4
	// num: 5
}
