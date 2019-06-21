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

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	var method, host, address string
	fmt.Sscanf(string(buf[:bytes.IndexByte(buf[:], '\n')]), "%s%s", &method, &host)

	hostPortUrl, err := url.Parse(host)
	if err != nil {
		fmt.Printf("parse err: %s\n", err.Error())
		return
	}

	if hostPortUrl.Opaque == "443" { // https访问
		address = hostPortUrl.Scheme + ":443"
	} else { // http访问
		if strings.Index(hostPortUrl.Host, ":") == -1 { // host不带端口，默认80
			address = hostPortUrl.Host + ":80"
		} else {
			address = hostPortUrl.Host
		}
	}

	fmt.Printf("[%d]: %s\n", n, string(buf[:n]))
	fmt.Println("method:", method, "host:", host, "address:", address)

	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Println("Dial:", err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(buf[:n])
	}
	//进行转发
	go io.Copy(server, conn)
	io.Copy(conn, server)
}
