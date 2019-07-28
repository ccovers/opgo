package main

import (
	"fmt"
	"time"
)

func main() {
	// defer语句调用一个函数，这个函数执行会推迟，直到外围的函数返回，
	// 或者外围函数运行到最后，或者相应的goroutine panic
	fmt.Println(*deferTest1())
	fmt.Println(deferTest2())

	// 函数值和函数参数被求值，但函数不会立即调用，实际函数调用却要等到最后
	foo("A")
	foo("B")

	//如果存在多个defer语句，最后的defer的函数的执行顺序与defer出现的顺序相反
	multiple()
}

func deferTest1() *int {
	v := 6
	defer func() {
		v = v + 1
	}()
	return &v
}

func deferTest2() int {
	v := 6
	defer func() {
		v = v + 1
	}()
	return v
}

func foo(name string) {
	defer trace(name)()
	fmt.Println("======分割线======")
}

func trace(funcName string) func() {
	start := time.Now()
	fmt.Printf("%s enter\n", funcName)
	return func() {
		fmt.Printf("%s exit (elapsed %s)\n", funcName, time.Since(start))
	}
}

func multiple() {
	printName := func(name string) {
		fmt.Println(name)
	}
	defer printName("A")
	defer printName("B")
	defer printName("C")
}
