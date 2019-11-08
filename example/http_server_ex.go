package main

import (
	"fmt"
	"io"
	"net/http"
)

func handle(res http.ResponseWriter, req *http.Request) {
	data := make([]byte, 1024)
	n, err := req.Body.Read(data)
	if err != nil && err != io.EOF {
		fmt.Printf("Read: %s\n", err.Error())
		return
	}

	fmt.Printf("Read success [%d]: %s\n", n, string(data))

	n, err = res.Write(data)
	if err != nil {
		fmt.Printf("Write: %s\n", err.Error())
		return
	}
	fmt.Printf("Write success [%d]\n", n)
}

func main() {
	http.HandleFunc("/v1", handle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("ListenAndServe: %s\n", err.Error())
		return
	}
}
