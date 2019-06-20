package main

import (
	"bufio"
	"fmt"
	//"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s <===> %s\n", conn.LocalAddr().String(), conn.RemoteAddr().String())

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	//defer c.Close()
	/*for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}*/

	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
