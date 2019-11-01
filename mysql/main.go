package main

import (
	"bytes"
	"fmt"
	"runtime"

	_ "github.com/siddontang/go-mysql/mysql"
)

func main() {
	//runtime.GOMAXPROCS(1)

	buf := call1()
	fmt.Printf("stack: [%s]\n", buf)
}

func call1() string {
	return call2()
}

func call2() string {
	return call3()
}

func call3() string {
	return Stack(0)
}

func Stack(skip int) string {
	buf := new(bytes.Buffer)

	callers := make([]uintptr, 32)
	n := runtime.Callers(skip, callers)
	frames := runtime.CallersFrames(callers[:n])
	for {
		if f, ok := frames.Next(); ok {
			fmt.Fprintf(buf, "%s\n\t%s:%d (0x%x)\n", f.Function, f.File, f.Line, f.PC)
		} else {
			break
		}
	}
	return buf.String()
}
