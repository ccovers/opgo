package main

import (
	"fmt"
)

func main() {
	fmt.Println(atoi(" -120 "))
}

func Atoi(numStr string) int {
	var flag bool = true
	var num int
	for _, v := range []byte(numStr) {
		if v >= '0' && v <= '9' {
			num = num*10 + int(v-'0')
		} else {
			if v == ' ' || v == '	' {
				if num == 0 {
					continue
				} else {
					break
				}
			} else if v == '-' {
				if num == 0 {
					flag = false
					continue
				}
			}
			return 0
		}
	}
	if !flag {
		num = -num
	}
	return num
}
