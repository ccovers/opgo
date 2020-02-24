package main

import (
	"fmt"
)

type MyOut struct {
}

func (w MyOut) Write(p []byte) (n int, err error) {
	return fmt.Printf("myout: %s", string(p))
}

var myOut MyOut

func main() {
	fmt.Fprintf(myOut, "xxx: %d\n", 10)
}
