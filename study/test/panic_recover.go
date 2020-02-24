package main

import (
	"fmt"
)

func test(n int, c chan int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		c <- n
	}()

	panic(fmt.Sprintf("panic: %d", n))
}

func main() {
	c := make(chan int, 0)

	for i := 0; i < 10; i++ {
		go test(i, c)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}
