package main

import (
	"fmt"
)

func main() {
	fmt.Println(itoa(-123000))
}

// 模拟
func itoa(num int) string {
	// return fmt.Sprintf("%d", num)
	var flag bool = true
	var vstr string
	if num < 0 {
		flag = false
		num = -num
	}

	for {
		o := num % 10
		vstr = string('0'+o) + vstr
		num = num / 10
		if num == 0 {
			break
		}
	}
	if !flag {
		vstr = "-" + vstr
	}
	return vstr
}
