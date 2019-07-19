package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	address := "localhost:8000"
	if len(os.Args) >= 2 {
		address = os.Args[1]
	}
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
	fmt.Println("over ...")
}

func mustCopy(dst io.Writer, src io.Reader) {
	/*if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}*/
	for {
		buf := make([]byte, 1024)
		_, err := src.Read(buf)
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		if string(buf) == "exit\n" {
			fmt.Println("eixt!")
			break
		}
		fmt.Println(string(buf))
		_, err = dst.Write(buf)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}
