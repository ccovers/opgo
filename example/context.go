package main

import (
	"context"
	"fmt"
	//"reflect"
	"time"
)

type MyCtx struct {
	context.Context
	NameMap map[int64]string
}

func (ctx MyCtx) Value(key interface{}) interface{} {
	var name string

	switch id := key.(type) {
	case int:
		name, _ = ctx.NameMap[int64(id)]
	case int64:
		name, _ = ctx.NameMap[id]
	}
	return name
}

/*
	参考自:
		https://zhuanlan.zhihu.com/p/58967892
*/
func main() {
	myCtx := MyCtx{
		Context: context.Background(),
		NameMap: map[int64]string{1: "一一一", 2: "二二二"},
	}

	ctx, cancel := context.WithTimeout(myCtx, 3*time.Second)
	ctx = context.WithValue(ctx, 3, "haha")
	go watch(ctx, "[监控1]")
	go watch(ctx, "[监控2]")
	go watch(ctx, "[监控3]")

	time.Sleep(3 * time.Second)
	fmt.Println("定时到达")
	time.Sleep(5 * time.Second)
	fmt.Println("ok, 通知停止")
	cancel()

	time.Sleep(5 * time.Second)

	name := ctx.Value(1)
	fmt.Println(name)
	name = ctx.Value(3)
	fmt.Println(name)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了")
			return
		default:
			fmt.Println(name, "goroutine监控中")
			time.Sleep(2 * time.Second)
		}
	}
}
