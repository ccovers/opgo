package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 通道
func main() {
	// struct{}类型的channel，它不能被写入任何数据，只有通过close()函数进行关闭操作，不占用任何内存！
	done := make(chan struct{})
	// ch := make(chan int)     // 无缓冲通道
	// ch1 := make(chan int, 0) // 无缓冲通道
	// ch2 := make(chan int, 3) // 缓冲通道

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		fmt.Println("===+")
		io.CopyN(os.Stdout, conn, 100)
		fmt.Println(done)
		close(done)
		fmt.Println("===++")
	}()
	fmt.Println("===")
	mustCopy(conn, os.Stdin)
	conn.Close()
	fmt.Println("===-")
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.CopyN(dst, src, 100); err != nil {
		fmt.Println(err)
	}
}
