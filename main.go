package main

import (
	"fmt"
	"unsafe"
)

type Inf struct {
	ID    int16
	IsOk1 int32
}

func main() {
	s := "h你"
	for i, v := range s {
		fmt.Printf("%c, %c\n", s[i], v)

	}
	fmt.Println("======")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	var x int
	fmt.Println(unsafe.Sizeof(x))
}
