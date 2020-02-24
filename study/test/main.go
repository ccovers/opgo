package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\r%s", "xx")
	time.Sleep(1 * time.Second)
	fmt.Printf("\r%s", "yy")
	time.Sleep(1 * time.Second)
	fmt.Printf("\r%s", "cc")
	time.Sleep(1 * time.Second)
}
