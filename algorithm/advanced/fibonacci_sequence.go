package main

import (
	"fmt"
)

/*
* 斐波那契数列
* 指的是这样一个数列：1、1、2、3、5、8、13、21、……在数学上，斐波那契数列以如下被以递归的方法定义：F0=0，F1=1，Fn=Fn-1+Fn-2（n>=2，n∈N*），用文字来说，就是斐波那契数列由 0 和 1 开始，之后的斐波那契数列系数就是由前面的两数相加。
 */

var cnt int

func main() {
	cnt = 0
	fmt.Println(fib(0), cnt)
	cnt = 0
	fmt.Println(fib(1), cnt)
	cnt = 0
	fmt.Println(fib(2), cnt)
	cnt = 0
	fmt.Println(fib(5), cnt)
	cnt = 0
	fmt.Println(fib(30), cnt)
	cnt = 0
	fmt.Println(fibEx(1, 1, 30), cnt)
}

// 斐波那契数
func fib(n int) int {
	cnt += 1
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 斐波那契数优化
// 通过观察旧的 fib 代码，我们发现一个问题，当给定一个数N时，需要先计算N-1 和N-2的情况，但是在计算N-1时同样要用计算N-2的情况，每个递归调用都触发另外两个递归调用，而这两个调用的任何一个又将调用另外连个递归调用，这样冗余计算的增量是非常快的。再计算 fib(5)时，需要先求fib(4)和fib(3)，而求fib(4)时又要求fib(3)和fib(2)以此类推这样就做了大量重复计算造成了不必要的浪费，如果我们能把fib(3),fib(2),fib(1)计算好的数据先存起来，下次需要计算时直接调用就可以省去大量的重复计算，因而有了下面的优化方式
func fibEx(a, b, n int) int {
	cnt += 1
	if n >= 2 {
		return fibEx(a+b, a, n-1)
	}
	return a
}
