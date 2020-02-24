package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}
