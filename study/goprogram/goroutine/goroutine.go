package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// 套接字接收过程
func socketRecv(conn net.Conn, wg *sync.WaitGroup) {
	// 创建一个接收的缓冲
	buff := make([]byte, 1024)

	// 不停地接收数据
	for {

		// 从套接字中读取数据
		_, err := conn.Read(buff)

		// 需要结束接收, 退出循环
		if err != nil {
			break
		}

	}

	// 函数已经结束, 发送通知
	wg.Done()
}

func main() {

	// 连接一个地址
	conn, err := net.Dial("tcp", "www.163.com:80")

	// 发生错误时打印错误退出
	if err != nil {
		fmt.Println(err)
		return
	}

	// 退出通道
	var wg sync.WaitGroup

	// 添加一个任务
	wg.Add(1)

	// 并发执行接收套接字
	go socketRecv(conn, &wg)

	// 在接收时, 等待1秒
	time.Sleep(3 * time.Second)

	// 主动关闭套接字
	conn.Close()

	// 等待goroutine退出完毕
	wg.Wait()
	fmt.Println("recv done")
}
