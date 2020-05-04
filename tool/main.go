package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp4", ":0")
	if err != nil {
		log.Fatalf("Failed to listen err=%v", err)
	}
	log.Printf("Starting http server on %s", lis.Addr().String())

	fmt.Println("port: ", int64(lis.Addr().(*net.TCPAddr).Port))
}
