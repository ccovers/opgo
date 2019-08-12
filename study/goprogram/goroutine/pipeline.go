package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		/*for {
			x, ok := <-naturals
			if !ok {
				break
			}
		}*/
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// printer
	for x := range squares {
		fmt.Println(x)
	}
	fmt.Println("over")
}