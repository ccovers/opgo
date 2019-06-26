package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	if conn == nil {
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	fmt.Printf("data:\n")
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			return
		}
		fmt.Printf("%s", string(buf))
		if n < 1024 {
			break
		}
	}
	fmt.Printf("\n\n")

	conn.Write(byte("{status:0}"))
}
