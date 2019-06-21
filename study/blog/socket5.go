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

type Socket5Proto struct {
	Ver      byte
	Nmethods byte
	Methods  *[]byte
}

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

	var buf [1024]byte
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	if buf[0] == 0x05 { // 只处理socket5协议
		// 客户端回应： socket服务不需要验证方式
		conn.Write([]byte(0x05, 0x00))
		n, err := conn.Read(buf)

		var host, port string
		switch buf[3] {

		}
	}
}
