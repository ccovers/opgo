package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("start listen ... 8080\n")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConn(conn)
	}
}

type Res struct {
	Status int
}

func handleConn(conn net.Conn) {
	if conn == nil {
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		fmt.Printf("data:\n")
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("Read err: %s\n", err.Error())
				return
			}
			fmt.Printf("%s", string(buf))
			if n < 1024 {
				break
			}
		}
		fmt.Printf("\n\n")

		res := Res{
			Status: 100,
		}
		bdata, err := json.Marshal(&res)
		if err != nil {
			fmt.Printf("Marshal err: %s\n", err.Error())
			return
		}

		_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 %d OK\nContent-Length:%d\n\r\n\r%v\n\r", http.StatusOK, len(bdata)+1, string(bdata))))
		if err != nil {
			fmt.Printf("Write err: %s\n", err.Error())
			return
		}
		break
	}
}
